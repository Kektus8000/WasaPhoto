package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Call the function that handle the User Login
	rt.router.POST("/session", rt.wrap(rt.DoLogin))

	// Call the function that handle the User Research
	rt.router.GET("/userProfile/:userID", rt.wrap(rt.GetUserProfile))

	//Call the function that handle the Set Username
	rt.router.POST("/user/:userID/username", rt.wrap(rt.SetMyUsername))

	//Call the function that handle the followUser
	rt.router.POST("/userProfile/:userID/following/", rt.wrap(rt.FollowUser))

	//Call the function that handle the unFollowUser
	rt.router.DELETE("/userProfile/:userID/following/user/:followerID", rt.wrap(rt.UnfollowUser))

	//Call the function that handle the getMyStream
	rt.router.GET("/userProfile/:userID/stream/", rt.wrap(rt.GetMyStream))

	//Call the function that handle the banUser
	rt.router.POST("/userProfile/:userID/banList/", rt.wrap(rt.BanUser))

	//Call the function that handle the unbanUser
	rt.router.DELETE("/userProfile/:userID/banList/user/:bannedID", rt.wrap(rt.UnBanUser))

	//Call the function that handle the uploadPhoto
	rt.router.POST("/userProfile/:userID/publishedPhotos/", rt.wrap(rt.UploadPhoto))

	//Call the function that handle the deletePhoto
	rt.router.DELETE("/userProfile/:userID/publishedPhotos/:photoID", rt.wrap(rt.DeletePhoto))

	//Call the function that handle the likePhoto
	rt.router.POST("/userProfile/:userID/stream/:photoID/like/", rt.wrap(rt.LikePhoto))

	//Call the function that handle the unlikePhoto
	rt.router.DELETE("/userProfile/:userID/stream/:photoID/like/", rt.wrap(rt.UnlikePhoto))

	//Call the function that handle the commentPhoto
	rt.router.POST("/userProfile/:userID/stream/:photoID/comments/", rt.wrap(rt.CommentPhoto))

	//Call the function that handle the uncommentPhoto
	rt.router.DELETE("/userProfile/:userID/stream/:photoID/comments/:commentID", rt.wrap(rt.UncommentPhoto))

	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
