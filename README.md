#Instagram-Basic-API
Developing a basic version of Instagram

The API's are developed using Go and MongoDB is used for storage.

Localhost port 8080 has been used.

This API system has following functions:

1) Create a User

  For creating the user, the url is http://localhost:8080/user

   The User model has following attributes:
          -User ID
          -Name
          -Email ID
          -Password --> Password has been encrypted and securely stored.
          
2) Get a user using id

  For getting the user using id, the url is http://localhost:8080/user/<user_id>   

3) Create a Post

    For creating a post, the url is http://localhost:8080/posts
    
    The post model has following attributes:
        -User ID (To know which user has posted the image)
        -Post ID
        -Caption
        -Image URL
        -Posted Time Stamp

4) Get a post using id

    For getting the post using id, the url is http://localhost:8080/posts/<post_id>

5) List of all posts of a user --> not working (fetching empty output)

    For getting the list of all the posts by a user, url assigned is http://localhost:8080/posts/user/<user_id>
