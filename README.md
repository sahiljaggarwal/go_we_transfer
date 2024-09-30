
# WeTransfer - File Transfer API

This project is a REST API built with Go Fiber for file uploading, URL generation, and download management. It includes Cloudinary integration for file storage, rate-limiting, and a cron job to delete expired files after 2 hours.

## Features

- Upload files up to **5MB** in size
- Files expire after **2 hours**
- Rate limiting to protect against abuse
- File storage on **Cloudinary**
- Automatic deletion of expired files via a **cron job**
- File URL shortener to share downloadable links

## Tech Stack

- **Go Fiber** for the web framework
- **PostgreSQL** for the database (via GORM)
- **Cloudinary** for file storage
- **Go Routine** for concurrency, improving performance and speed
- **Channels** for communication between goroutines
- **Cron Jobs** for periodic tasks
- **Rate Limiting**, **Helmet** for security
- **Fiber Middlewares** (CORS, Logger, Cache, Monitor)

---

## Setup Instructions

### 1. Clone the repository

To get started, clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/super_crud.git
cd super_crud
```

### 2. Set up the environment variables

Create a `.env` file in the root directory with the following content:

```env
# Cloudinary API credentials
CLOUDINARY_CLOUD_NAME=dh8g2mhqs
CLOUDINARY_API_KEY=599757451293173
CLOUDINARY_API_SECRET=Ep1UY_glzDStYFoNtqnPhJnCiO8

# Server configuration
PORT=3000
HOST=http://localhost

# PostgreSQL Database credentials
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=ruka3030&
DB_NAME=super_crud
DB_PORT=5432
```

### 3. Build and Run the Project

You can build and run the project using the provided script:

```bash
chmod +x run_build.sh
./run_build.sh
```

This will:
- Create a `dist` folder
- Build the Go application into the `dist` folder
- Run the application

### 4. Access the API

Once the server is running, you can access the API on `http://localhost:3000`.

---

## API Endpoints

### File Upload

- **POST** `/file/upload`  
  Upload a file up to 5MB in size. The uploaded file will generate a URL.

### Get File by Short ID

- **GET** `/file/:id`  
  Retrieve a file using its short ID and download it.

### Get All Files

- **GET** `/files`  
  Fetch all uploaded files.

### Delete File by ID

- **DELETE** `/file/:id`  
  Permanently delete a file by its short ID.

---

## Cron Job

The cron job runs every hour and deletes files that have expired after 2 hours.

---

## Go Routines and Channels

To enhance the performance of the API, Go routines are utilized for handling concurrent requests, allowing the application to process multiple uploads simultaneously. Channels facilitate communication between goroutines, ensuring efficient data handling and reducing latency during file processing.

---

## Middleware Used

- **Rate Limiter**: Limits the number of requests to protect the API. Max 20 requests per 30 seconds.
- **Helmet**: Adds security headers to requests.
- **CORS**: Allows cross-origin requests.
- **Cache**: Caches GET requests for 30 seconds.
- **Logger**: Logs all requests.
- **Monitor**: Provides metrics for the server at `/metrics`.

---

## Rate Limiting

The API limits the requests using the following configuration:
- **Max**: 20 requests
- **Expiration**: 30 seconds

---

## Cloudinary

Files are uploaded and stored on **Cloudinary**. The API generates a shortened URL that users can share with others for downloading.

---

## Inspiration

This project is inspired by [WeTransfer](https://wetransfer.com/), offering a simple file upload and sharing service.

---

## License

This project is licensed under the MIT License.
