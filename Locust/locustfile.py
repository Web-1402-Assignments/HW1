from locust import HttpUser, task, between

class HelloWorldUser(HttpUser):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.nonce = ""
        self.servernonce = ""
        self.auth_key = ""

    wait_time = between(1, 5)
    @task
    def req_pq(self):
        with self.client.post('/auth/req_pq/'
                         , 
                         json={
                            "ID": 2,
                         }
                    ) as response:
            self.nonce = response.json["nonce"]
            self.servernonce = response.json["serverNonce"]
          
    
    @task
    def req_dh(self):
       with self.client.post('/auth/req_DH_params/'
                         , 
                         json={
                            'NONCE': self.nonce,
                            'SERVER_NONCE': self.servernonce,
                            'ID': 4,
                            'KEY': 6,
                         }
                    ) as response:
            self.auth_key = response.json["key"]  
            

    @task
    def get_users(self):
        self.client.post(
            '/biz/getUsers/',
            json={
                'AUTH_KEY': self.auth_key,
                'USER_ID': '1111',
                'MESSAGE_ID': 60,
            }
        )
    
    @task
    def get_users_with_injection(self):
        self.client.post(
            '/biz/WithInjection/',
            json={
                'AUTH_KEY': self.auth_key,
                'USER_ID': '1111 OR 1=1 --',
                'MESSAGE_ID': 50,
            }
        )