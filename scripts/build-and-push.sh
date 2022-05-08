PORT=5000
BACKEND_IMAGE_TAG="localhost:$PORT/simple-go-backend"
FRONTEND_IMAGE_TAG="localhost:$PORT/simple-react-frontend"


docker build backend -t "$BACKEND_IMAGE_TAG"
docker push "$BACKEND_IMAGE_TAG"


docker build frontend -t "$FRONTEND_IMAGE_TAG"
docker push "$FRONTEND_IMAGE_TAG"