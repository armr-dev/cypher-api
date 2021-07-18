# cypher-api
![image](https://img.shields.io/github/go-mod/go-version/armr-dev/cypher-api)
[![image](https://img.shields.io/badge/repository-cypher--front-orange)](https://github.com/armr-dev/cypher-front)

`cypher-api` is an API that cypher/decipher messages using one of five block cypher algorithms.
It's used mainly as part of the Cypher project, that is composed by `cypher-api` and `cypher-front`. You can check out `cypher-front` [here](https://github.com/armr-dev/cypher-front).

:warning: ***This should not be used in production because the key for the algorithms are hardcoded in the application, for now.***

This api is hosted using Heroku and you can check it out [here](https://cypher-api-go.herokuapp.com).

This project was done as part of a course in university.

### Usage
There are 2 main endpoints in the API: `/cypher` and `/decipher`.

To cypher and decipher messages, all you have to do is make a POST request to the api with a body in the following format:
```ts
  {
    text: string,
    algorithm: "aes" | "des" | "3des" | "blowfish" | "idea"
  }
```

The API will then respond with a JSON in the following format:
```ts
  {
    data: {
      text: string
    }
  }
```
where `text` is the cyphered/deciphered string.

### ToDo
 - [ ] Add key param for algorithms;
