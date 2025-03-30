# 🏨 Hotel Rental App (W.I.P) 🚀

## 🌟 Overview

The **Hotel Rental App** is a web-based platform that allows users to browse, book, and manage hotel reservations seamlessly. It provides real-time availability, secure payments, and an intuitive user experience for both travelers and hotel managers. 🏝️✨

## 🔥 Features

### 🎯 User Features

- 🔍 **Hotel Search & Filtering**: Find the perfect stay by filtering hotels based on location, price, ratings, and amenities. 🏙️
- ⏳ **Real-time Booking System**: Instantly check availability and reserve rooms with one click! ✅
- 🌟 **User Reviews & Ratings**: Read and leave reviews to help future travelers make informed decisions. 📝✨
- 👤 **User Authentication & Profiles**: Register, log in, and manage bookings with a secure and personalized account. 🔑

### 🛠️ Admin Features

- 🏢 **Hotel Management**: Add, update, or remove hotel listings with images, descriptions, and prices. 📸📌
- 📊 **Booking Dashboard**: Get a complete overview of all reservations, making hotel operations smooth and efficient. 🏆
- 💰 **Revenue Insights**: Track earnings in real time and optimize pricing for better business strategy. 📈💡

## ⚙️ Tech Stack

- 🏎 **Backend**: Golang ensures high performance and concurrency for seamless real-time booking operations. ⚡
- 🔗 **Interservice Communication**: gRPC & Redpanda enable lightning-fast, low-latency communication between services. 🚀
- 🔄 **Workflow Management**: Temporal.io makes booking workflows resilient, ensuring fault tolerance. 🛡️
- 🗄️ **Database & Search**: MongoDB (Primary Data Source) for structured hotel and booking data, Memgraph for complex graph-based queries, and Elasticsearch for fast, intelligent search. 🔍📊
- 🏗️ **Containerization & Orchestration**: Docker & Kubernetes provide seamless deployment and scalability. 📦⚙️
- 🔐 **Service Discovery & Security**: Consul ensures reliable service discovery, while Vault securely manages sensitive credentials. 🔑
- 📜 **API Documentation**: Swagger makes API endpoints clear and easy to use for developers. 🛠️
- ✅ **CI/CD & Testing**: GitHub Actions automates testing and deployment, ensuring smooth development cycles. 🔄
- 📩 **Notifications**: SMTP sends automated emails for confirmations, updates, and alerts. 📬
- 🗺️ **Google Maps Integration** – Enhanced search with location-based recommendations. 📍

## 🚀 Installation & Setup

### 🔧 Backend Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/hotel-rental-app.git
cd hotel-rental-app/backend

# Build and run the backend
docker-compose up --build
```

## 🔗 API Endpoints

| 🔹 Method | 🔸 Endpoint         | 📌 Description         |
| --------- | ------------------- | ---------------------- |
| GET       | /hotels             | Get list of hotels     |
| GET       | /hotels/{id}        | Get details of a hotel |
| POST      | /bookings           | Create a new booking   |
| GET       | /bookings/{user_id} | Get user bookings      |
| POST      | /auth/signup        | Register a new user    |
| POST      | /auth/login         | Authenticate user      |

**More incoming**

## 🚀 Future Enhancements

- 🤖 **AI-based Hotel Recommendations** – Personalized suggestions based on user preferences. 🎯
- 💲 **Dynamic Pricing Based on Demand** – Adaptive pricing strategy for better revenue. 📈
- 💬 **In-app Customer Support Chatbot** – Instant support for users via AI chat. 🤖💡
- 💳 **Secure Payment Integration**: Pay with confidence using Stripe and PayPal, ensuring encrypted and safe transactions. 🔐
