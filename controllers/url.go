package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"url-shortener/models"
	"url-shortener/utils"
)

func CreateURL(c *gin.Context) {
	var urlInput struct {
		OriginalURL string `json:"original_url"`
	}

	if err := c.ShouldBindJSON(&urlInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := primitive.ObjectIDFromHex(c.GetString("userID"))

	// Generate short code
	shortCode := utils.GenerateShortCode()

	// Create URL
	newURL := models.URL{
		OriginalURL: urlInput.OriginalURL,
		ShortCode:   shortCode,
		UserID:      userID,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(15 * 24 * time.Hour), // 15 days
		ClickCount:  0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := utils.DB.Collection("urls").InsertOne(ctx, newURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create URL"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_url":    "http://" + c.Request.Host + "/" + shortCode,
		"original_url": urlInput.OriginalURL,
		"expires_at":   newURL.ExpiresAt,
	})
}

func GetUserURLs(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userID"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := utils.DB.Collection("urls").Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch URLs"})
		return
	}
	defer cursor.Close(ctx)

	var urls []models.URL
	if err = cursor.All(ctx, &urls); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode URLs"})
		return
	}

	c.JSON(http.StatusOK, urls)
}

func GetURL(c *gin.Context) {
	urlID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(urlID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL ID"})
		return
	}

	userID, _ := primitive.ObjectIDFromHex(c.GetString("userID"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var url models.URL
	err = utils.DB.Collection("urls").FindOne(ctx, bson.M{"_id": objID, "user_id": userID}).Decode(&url)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch URL"})
		return
	}

	c.JSON(http.StatusOK, url)
}

func DeleteURL(c *gin.Context) {
	urlID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(urlID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL ID"})
		return
	}

	userID, _ := primitive.ObjectIDFromHex(c.GetString("userID"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := utils.DB.Collection("urls").DeleteOne(ctx, bson.M{"_id": objID, "user_id": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete URL"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL deleted successfully"})
}

func RedirectURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var url models.URL
	err := utils.DB.Collection("urls").FindOneAndUpdate(
		ctx,
		bson.M{"short_code": shortCode, "expires_at": bson.M{"$gt": time.Now()}},
		bson.M{"$inc": bson.M{"click_count": 1}},
	).Decode(&url)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found or expired"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not redirect"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

func CleanExpiredURLs() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := utils.DB.Collection("urls").DeleteMany(ctx, bson.M{"expires_at": bson.M{"$lt": time.Now()}})
	if err != nil {
		log.Println("Error cleaning expired URLs:", err)
	}
}
