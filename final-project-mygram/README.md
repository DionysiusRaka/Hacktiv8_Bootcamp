# MyGram - Photo Sharing and Commenting App **(API ONLY)**

## Description

MyGram is a captivating photo-sharing and commenting application designed for the final project. Seamlessly share your favorite photos and engage in meaningful conversations with the MyGram community.

## Key Features

- **Photo Sharing:** Upload and showcase your moments effortlessly.
- **Commenting System:** Foster connections by leaving comments on others' photos.
- **User Registration:** Create your MyGram account for a personalized experience.
- **Authentication:** Ensure a secure and enjoyable user experience.
- **User Profiles:** Customize your profile and explore a diverse range of photos.

## Endpoints

### User Section

- `POST /users/register`: Register a new user.
- `POST /users/login`: Log in to MyGram. *This process gives you a bearer token; use it for authorization.*

### Photo Section

- `POST /photos/`: Create a new photo (authentication required).
- `GET /photos/`: Get all photos.
- `GET /photos/:photoId`: Get a specific photo.
- `PUT /photos/:photoId`: Update a specific photo. *Authentication and authorization required.*
- `DELETE /photos/:photoId`: Delete a specific photo. *Authentication and authorization required.*

### Comment Section

- `GET /comments/`: Get all comments. *Authentication required.*
- `GET /comments/:commentId`: Get a specific comment. *Authentication required.*
- `POST /comments/:photoId`: Create a new comment. *Authentication required.*
- `PUT /comments/:commentId`: Update a specific comment. *Authentication and authorization required.*
- `DELETE /comments/:commentId`: Delete a specific comment. *Authentication and authorization required.*

### Social Media Section

- `POST /socmed/`: Create a new social media post. *Authentication required.*
- `GET /socmed/`: Get all social media posts.
- `GET /socmed/:socMedId`: Get a specific social media post.
- `PUT /socmed/:socMedId`: Update a specific social media post. *Authentication and authorization required.*
- `DELETE /socmed/:socMedId`: Delete a specific social media post. *Authentication and authorization required.*

## How to Use
> **REMEMBER TO DO THIS IN POSTMAN USING THE ENDPOINTS GIVEN ABOVE!!!**
> 
1. **Register:** Sign up for MyGram by creating an account.
2. **Explore Photos:** Discover a variety of photos shared by other users.
3. **Share Moments:** Upload and share your own photos to create a personalized feed.
4. **Leave Comments:** Engage with the community by leaving comments on photos.
5. **Connect with Others:** Follow users, build connections, and create a network of MyGram friends.

MyGram is more than just a photo-sharing app; it's a vibrant community where visual storytelling and meaningful interactions thrive. Join MyGram today and start sharing your world with others!
