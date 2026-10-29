import time

def fetch_user():
    time.sleep(2)
    print("User loaded")

def fetch_products():
    time.sleep(2)
    print("Products loaded")

def fetch_notifications():
    time.sleep(2)
    print("Notifications loaded")

start = time.time()

fetch_user()
fetch_products()
fetch_notifications()

print("Total:", round(time.time() - start, 2), "seconds")