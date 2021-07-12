
# Jeopardy! Question Search REST API

A simple REST API that allows a user to upload a CSV of Jeopardy questions.

After using a POST request to upload these CSV users can then search via monetary value of the question, category of the question, round of Jeopardy in which the question was asked, and finally by round and category. All of these are provided through GET requests mapped within the API. Finally the user is able to delete questions by ID by using a DELETE request, and add a question using the body of a POST request and a simple Question object.



## API Reference

#### Upload Questions

```http
  POST /upload
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `file` | `CSV file` | **Required**. A CSV file of Jeopardy questions.|

#### Get number of questions from start questions

```http
  GET /questions/{num}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| Num      | Integer |Returns a desired number of questions.
#### Get question by ID

```http
  GET /question/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| ID      | Integer |Returns the question associated with a given ID.

#### Get question by value

```http
  GET /questionsByValue/{value}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| Value      | Integer |Returns a list of all questions worth a given value.

#### Get question by category

```http
  GET /questionsByCategory/{category}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| category      | String |Returns a list of all questions within a given categroy.

#### Get question by round

```http
  GET /questionsByRound/{round}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| Round      | String |Returns a list of all questions asked within a given round.

#### Get question by round and category

```http
  GET /questionsByRoundAndCategory/{round}/{category}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| round / category      | String / String|Returns a list of all questions within a given category.

#### Delete question by ID

```http
  DELETE /deleteQuestion/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| ID      | Long |Deletes the question at the given ID.

#### Add a question

```http
  POST /addQuestion/{shownumber}/{round}/{category}/{value}/{question}/{answer}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| shownumber, round, category, value, question, answer     | Question |Adds a question to the database when given the parameters.

### Question Object:

```
    ShowNumber int    
	Round      string 
	Category   string 
	Value      int    
	Question   string 
	Answer     string 
```

  
## Tech Stack

**Client:** Golang, Gorilla MUX, GORM

  
