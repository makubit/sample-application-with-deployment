# This file contains demo of deployment to existing EKS cluster

1. Start with building project and running tests
![make build](img/make%20build.png)
2. Build Docker image
![make docker-build](img/make%20docker-build.png)
3. Push Docker image
![make docker-push](img/make%20docker-push.png)
4. Run helm deployment
![make helm-deploy](img/make%20helm-deploy.png)
5. Check if all resources are running
![get po](img/get%20po.png)
![get secret](img/get%20secret.png)
6. Check if API is responding (using port-forward)
![api 200](img/api%20200.png)
7. Check if object exists in database
![mongo](img/mongo.png)

I hope you enjoyed this demo and whole project :)

