package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/temog/blog/backend/models/article"
	"github.com/temog/blog/backend/models/image"
	"github.com/temog/blog/backend/models/tag"
	"github.com/temog/blog/backend/models/user"
	"github.com/temog/blog/backend/util"
)

const Base = "/api/"

func main() {

	// .env 読み込み
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	router := gin.Default()

	// Cross Origin Resource Sharing
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("ALLOW_ORIGINS")}
	config.AddAllowHeaders("authorization")
	router.Use(cors.New(config))

	/**** api/user ****/
	user := router.Group(Base + "user/")
	{
		// api/user/auth
		user.OPTIONS("auth", Util.ResponseOption)
		user.POST("auth", func(c *gin.Context) {

			input := User.InputAuth{}
			if err := c.ShouldBind(&input); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			status, err, token := User.Auth(input)
			c.JSON(http.StatusOK, gin.H{
				"status": status,
				"error":  getError(err),
				"token":  token,
			})
		})
	}

	/**** api/tag ****/
	tag := router.Group(Base + "tag/")
	{
		// api/tag/add
		tag.OPTIONS("add", Util.ResponseOption)
		tag.POST("add", func(c *gin.Context) {

			if err := User.ValidateToken(c.Request.Header.Get("Authorization")); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			input := Tag.InputAdd{}
			if err := c.ShouldBind(&input); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			status, err, objectId := Tag.Add(input)
			c.JSON(http.StatusOK, gin.H{
				"status":   status,
				"objectId": objectId,
				"error":    getError(err),
			})
		})

		// api/tag/update
		tag.OPTIONS("update", Util.ResponseOption)
		tag.POST("update", func(c *gin.Context) {

			if err := User.ValidateToken(c.Request.Header.Get("Authorization")); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			input := Tag.InputUpdate{}
			if err := c.ShouldBind(&input); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			status, err := Tag.Update(input)
			c.JSON(http.StatusOK, gin.H{
				"status": status,
				"error":  getError(err),
			})
		})

		// api/tag/getAll
		tag.OPTIONS("getAll", Util.ResponseOption)
		tag.GET("getAll", func(c *gin.Context) {

			status, err, tags := Tag.GetAll()

			c.JSON(http.StatusOK, gin.H{
				"status": status,
				"error":  getError(err),
				"tags":   tags,
			})
		})
	}

	/**** api/article ****/
	article := router.Group(Base + "article/")
	{
		// api/article/add
		article.OPTIONS("add", Util.ResponseOption)
		article.POST("add", func(c *gin.Context) {

			if err := User.ValidateToken(c.Request.Header.Get("Authorization")); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			input := Article.InputAdd{}
			if err := c.ShouldBind(&input); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			status, err, objectId := Article.Add(input)
			c.JSON(http.StatusOK, gin.H{
				"status": status,
				"error":  getError(err),
				"id":     objectId,
			})
		})

		// api/article/edit
		article.OPTIONS("edit", Util.ResponseOption)
		article.POST("edit", func(c *gin.Context) {

			if err := User.ValidateToken(c.Request.Header.Get("Authorization")); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			input := Article.InputEdit{}
			if err := c.ShouldBind(&input); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			status, err := Article.Edit(input)
			c.JSON(http.StatusOK, gin.H{
				"status": status,
				"error":  getError(err),
			})
		})

		// api/article/getPopularArticles
		article.OPTIONS("getPopularArticles", Util.ResponseOption)
		article.GET("getPopularArticles", func(c *gin.Context) {

			status, err, articles := Article.GetPopularArticles()
			c.JSON(http.StatusOK, gin.H{
				"status":   status,
				"error":    getError(err),
				"articles": articles,
			})
		})

		// api/article/find/:articleId
		article.OPTIONS("find/:articleId", Util.ResponseOption)
		article.GET("find/:articleId", func(c *gin.Context) {

			if err := User.ValidateToken(c.Request.Header.Get("Authorization")); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			articleId := c.Param("articleId")
			if articleId == "" {
				Util.ResponseValidationError(c, errors.New("articleId required"))
			}

			status, err, article := Article.FindId(articleId)

			c.JSON(http.StatusOK, gin.H{
				"status":  status,
				"error":   getError(err),
				"article": article,
			})
		})

		// api/article/alias/:alias
		article.OPTIONS("alias/:alias", Util.ResponseOption)
		article.GET("alias/:alias", func(c *gin.Context) {

			alias := c.Param("alias")
			if alias == "" {
				Util.ResponseValidationError(c, errors.New("alias required"))
			}

			status, err, article := Article.FindAlias(alias)
			c.JSON(http.StatusOK, gin.H{
				"status":  status,
				"error":   getError(err),
				"article": article,
			})
		})

		// api/article/search
		article.OPTIONS("search", Util.ResponseOption)
		article.POST("search", func(c *gin.Context) {

			input := Article.InputSearch{}
			if err := c.ShouldBind(&input); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			// admin用
			if input.Status != "publish" {
				if err := User.ValidateToken(c.Request.Header.Get("Authorization")); err != nil {
					Util.ResponseValidationError(c, err)
					return
				}
			}

			status, err, totalPage, articles := Article.Search(input)
			c.JSON(http.StatusOK, gin.H{
				"status":    status,
				"error":     getError(err),
				"totalPage": totalPage,
				"articles":  articles,
			})
		})
	}

	/**** api/image ****/
	image := router.Group(Base + "image/")
	{
		// api/image/add
		image.OPTIONS("add", Util.ResponseOption)
		image.POST("add", func(c *gin.Context) {

			if err := User.ValidateToken(c.Request.Header.Get("Authorization")); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			input := Image.InputAdd{}
			if err := c.ShouldBind(&input); err != nil {
				Util.ResponseValidationError(c, err)
				return
			}

			status, err, fileName := Image.Add(input)
			c.JSON(http.StatusOK, gin.H{
				"status":   status,
				"fileName": fileName,
				"error":    getError(err),
			})
		})
	}

	// 指定なしだと 8080
	// router.Run(":8080")
	router.Run(":8081")
}

func getError(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
