package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Metodi GET
	rt.router.GET("/user/:userID", rt.wrap(rt.GetUser))
	rt.router.GET("/userProfile/:userID", rt.wrap(rt.GetUserProfile))
	rt.router.GET("/userProfile/:userID/stream/", rt.wrap(rt.GetMyStream))
	rt.router.GET("/userProfile/:userID/publishedPhotos/:photoID", rt.wrap(rt.RecoverPhoto))

	// Metodi POST
	rt.router.POST("/session", rt.wrap(rt.DoLogin))
	rt.router.POST("/userProfile/:userID/publishedPhotos/", rt.wrap(rt.UploadPhoto))
	rt.router.POST("/userProfile/:userID/stream/:photoID/comments/", rt.wrap(rt.CommentPhoto))

	// Metodi PUT
	rt.router.PUT("/user/:userID/username", rt.wrap(rt.SetMyUsername))
	rt.router.PUT("/user/", rt.wrap(rt.SearchUsers))
	rt.router.PUT("/userProfile/:userID/banList/:bannedID", rt.wrap(rt.BanUser))
	rt.router.PUT("/userProfile/:userID/stream/:photoID/likes/", rt.wrap(rt.LikePhoto))
	rt.router.PUT("/userProfile/:userID/following/:followerID", rt.wrap(rt.FollowUser))

	// Metodi DELETE
	rt.router.DELETE("/userProfile/:userID/following/:followerID", rt.wrap(rt.UnfollowUser))
	rt.router.DELETE("/userProfile/:userID/banList/:bannedID", rt.wrap(rt.UnBanUser))
	rt.router.DELETE("/userProfile/:userID/publishedPhotos/:photoID", rt.wrap(rt.DeletePhoto))
	rt.router.DELETE("/userProfile/:userID/stream/:photoID/likes/", rt.wrap(rt.UnlikePhoto))
	rt.router.DELETE("/userProfile/:userID/stream/:photoID/comments/:commentID", rt.wrap(rt.UncommentPhoto))

	// Special routes
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
