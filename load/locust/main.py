from locust import HttpUser, task


class HelloWorldUser(HttpUser):
    @task
    def hello_world(self):
        self.client.get("/hello")

    @task
    def health(self):
        self.client.get("/health")
