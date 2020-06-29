from locust import task, between
from locust.contrib.fasthttp import FastHttpUser


class RollingNodeUser(FastHttpUser):
    wait_time = between(5, 9)

    @task
    def spam_head(self):
        self.client.get("/chains/main/blocks/head")

    def on_start(self):
        pass
