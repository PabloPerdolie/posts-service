# Posts Service 

Posts Service is a GraphQL-based API service for managing posts and comments.

## Features

- **Create Posts**: Users can create new posts with titles, content, and enable or disable comments.
- **Retrieve Posts**: Retrieve a list of posts with details including their comments.
- **Manage Comments**: Enable or disable comments on posts, view all comments associated with a post.
- **Subscribe to comments**: Subscribe to comments on a post and when you add a new comment to this post, you will receive an alert

## Technologies Used

- **GraphQL**: API for querying and mutating data.
- **Golang**: Backend service implemented in Go.
- **PostgreSQL**: Relational database for persistent data storage.
- **Docker**: Containerization for easy deployment and scalability.

## Setup Instructions

Follow these steps to set up and run the Posts Service locally:

### Clone the repository

```bash
git clone https://github.com/PabloPerdolie/posts-service.git
cd posts-service
```

### Launching the application

#### The application will be launching with in-memory storage:
```bash
make in_memory
```

#### The application will be launching with PostgreSQL storage:
```bash
make postgresql
```
The service will start running on http://localhost:8080/.

#### To clean up:
```bash
make clean
```

## Access GraphQL Playground

Open your web browser and navigate to http://localhost:8080/ to interact with the GraphQL Playground.

## GraphQL Operations

Use GraphQL queries and mutations to interact with the service:

#### Create Post

```graphql
mutation {
    createPost(title: "New Post", content: "Content of the post", commentsEnabled: true) {
        id
        title
        content
        commentsEnabled
    }
}
```

#### Create Comment

```graphql
mutation CreateComment {
    createComment(
        postId: "postId"
        content: "Content of the comment"
        parentId: "Parent id of comment you reply")
    {
        id
        postId
        content
        createdAt
        parentId
    }
}
```

#### Get Post with Comments

```graphql
query {
    post(id: "postId") {
        id
        title
        content
        commentsEnabled
    }
    comments(postId: "postId") {
        id
        content
        createdAt
        children {
            id
            content
            createdAt
            children {
                id
                content
                createdAt
                # etc
            }
        }
    }
}
```
#### Subscribe to Post's comments

_Note: after one alert you can for wait for the next cause this app used socket connection for transport_

```graphql
subscription CommentAdded {
    commentAdded(postId: "postId") {
        id
        postId
        parentId
        content
        createdAt
    }
}
```

#### Get all Posts

```graphql
query {
    posts {
        id
        title
        content
        commentsEnabled
    }
}
```

#### Manage Comments on a Post

```graphql
mutation {
    manageComments(postID: "postId", enable: true) {
        id
        title
        content
        commentsEnabled
    }
}
```

#### Get Comments with pagination

_Note: LIMIT is intended to set the number of records to be returned. OFFSET specifies to skip the specified number of rows before starting to output rows_

```graphql
query {
    comments(postId: "postId", limit: 0, offset: 10) {
        id
        content
        createdAt
        children {
            id
            content
            createdAt
            children {
                id
                content
                createdAt
                # etc
            }
        }
    }
}
```

<center>Thanks for checking out my service.</center>
