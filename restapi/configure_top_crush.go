// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	log "github.com/sirupsen/logrus"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/capella/topcrush/restapi/operations"
	"github.com/capella/topcrush/restapi/operations/chat"
	"github.com/capella/topcrush/restapi/operations/matchs"
	"github.com/capella/topcrush/restapi/operations/slide"
	"github.com/capella/topcrush/restapi/operations/user"

	models "github.com/capella/topcrush/models"
)

//go:generate swagger generate server --target ../../top-crush --name TopCrush --spec ../swagger.yml --principal models.Principal

func configureFlags(api *operations.TopCrushAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TopCrushAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Token" header is set
	api.APIKeyAuthAuth = func(token string) (*models.Principal, error) {
		return nil, errors.NotImplemented("api key auth (ApiKeyAuth) Token from header param [Token] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.MatchsDeleteMatchsDeleteHandler == nil {
		api.MatchsDeleteMatchsDeleteHandler = matchs.DeleteMatchsDeleteHandlerFunc(func(params matchs.DeleteMatchsDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation matchs.DeleteMatchsDelete has not yet been implemented")
		})
	}
	if api.ChatGetChatMessagesIDMessageIndexHandler == nil {
		api.ChatGetChatMessagesIDMessageIndexHandler = chat.GetChatMessagesIDMessageIndexHandlerFunc(func(params chat.GetChatMessagesIDMessageIndexParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation chat.GetChatMessagesIDMessageIndex has not yet been implemented")
		})
	}
	if api.MatchsGetMatchsAllHandler == nil {
		api.MatchsGetMatchsAllHandler = matchs.GetMatchsAllHandlerFunc(func(params matchs.GetMatchsAllParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation matchs.GetMatchsAll has not yet been implemented")
		})
	}
	if api.MatchsGetMatchsLikesHandler == nil {
		api.MatchsGetMatchsLikesHandler = matchs.GetMatchsLikesHandlerFunc(func(params matchs.GetMatchsLikesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation matchs.GetMatchsLikes has not yet been implemented")
		})
	}
	if api.SlideGetSlideUsersHandler == nil {
		api.SlideGetSlideUsersHandler = slide.GetSlideUsersHandlerFunc(func(params slide.GetSlideUsersParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation slide.GetSlideUsers has not yet been implemented")
		})
	}
	if api.UserGetUserHandler == nil {
		api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		})
	}
	if api.UserGetUserUploadHandler == nil {
		api.UserGetUserUploadHandler = user.GetUserUploadHandlerFunc(func(params user.GetUserUploadParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserIDUpload has not yet been implemented")
		})
	}
	if api.UserPostUserHandler == nil {
		api.UserPostUserHandler = user.PostUserHandlerFunc(func(params user.PostUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.PostUser has not yet been implemented")
		})
	}
	if api.ChatPutChatMessagesIDHandler == nil {
		api.ChatPutChatMessagesIDHandler = chat.PutChatMessagesIDHandlerFunc(func(params chat.PutChatMessagesIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation chat.PutChatMessagesID has not yet been implemented")
		})
	}
	if api.SlidePutSlideBoastHandler == nil {
		api.SlidePutSlideBoastHandler = slide.PutSlideBoastHandlerFunc(func(params slide.PutSlideBoastParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation slide.PutSlideBoast has not yet been implemented")
		})
	}
	if api.SlidePutSlideLikeHandler == nil {
		api.SlidePutSlideLikeHandler = slide.PutSlideLikeHandlerFunc(func(params slide.PutSlideLikeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation slide.PutSlideLike has not yet been implemented")
		})
	}
	if api.SlidePutSlideSuperlikeHandler == nil {
		api.SlidePutSlideSuperlikeHandler = slide.PutSlideSuperlikeHandlerFunc(func(params slide.PutSlideSuperlikeParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation slide.PutSlideSuperlike has not yet been implemented")
		})
	}
	if api.UserPutUserIDHandler == nil {
		api.UserPutUserIDHandler = user.PutUserIDHandlerFunc(func(params user.PutUserIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.PutUserID has not yet been implemented")
		})
	}
	if api.UserPutUserIDLocationHandler == nil {
		api.UserPutUserIDLocationHandler = user.PutUserIDLocationHandlerFunc(func(params user.PutUserIDLocationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.PutUserIDLocation has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
