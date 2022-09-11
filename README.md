# Go Programming Challenge

Hello 👋
This is Onqlave's challenge for potential new engineering team members.

## Instructions

Fork this repository, and implement the scenario described below.

* You *shall* complete this challenge using **Go language**,
* You *should* try to show your development process to present a **production-ready** code,
* Please, describe your approach, your technical choices, including architectural, and anything you want us to know in a results.md file,

## The Scenario

Your mission is to implement a single endpoint that will be integrated in a larger **microservice** architecture.

* the endpoint is exposed at `/api/v1/applications`
* the endpoint allows to manage data from a local database (see explanations below)
* the query may takes parameters, to filter the results. You are expected to propose any kind of query params you may find useful to provide a rich user experience,
* the endpoint returns a JSON response with an array of matching results
* application should be delivered as a docker + docker-compose so we cna start it locally without any dependency

## The database

The initial database is provided in a file [db.json](db.json).
You are expected to integrate the data in a mongodb local instance (docker), and explain how to install, launch and populate the database. This database **must** be used by your code.
The schema of the provided data is :

```
{
  "type": "array",
  "items": {
    "type": "object"
    {
      "name": {
        "type": "string",
        "minLength": 1
      },
      "description": {
        "type": "string",
	     "minLength": 1,
	     "maxLength": 150
      },
      "enabled": {
        "type": "boolean"
      },
      "type": {
        "type": "string"
      }
    }
  }
}
```

Example:

```
[
  {
    "text": "Accounting application.",
    "description": "1.",
    "enabled": true,
    "type": "web"
  },
  {
    "text": "Human resource application.",
    "description": "1.",
    "enabled": true,
    "type": "service"
  }
]
```

## How we review ?

Your application will be reviewed by at one of our team member. We do take into consideration your experience level.

**We value quality over feature-completeness**. The goal of this code sample is to help us identify what you consider **production-ready** code. You should consider this code ready for final review with your colleague, i.e. this would be the last step before deploying to production.

The aspects of your code we will assess include:

* **Architecture**: how clean is the separation between the logical layers ? does the code make use of design patterns ? Are the RESTful principles followed ?
* **Clarity**: does the documentation (result.md) clearly and concisely explains the problem and solution? Are technical tradeoffs explained ?
* **Correctness**: does the code do what was asked ? If there is anything missing, does the result.md explain why it is missing?
* **Code quality**: is the code simple, easy to understand, and maintainable ? Are there any code smells or other red flags ? Does it follows coding principles such as the single responsibility principle ? Is the coding style consistent with the language's guidelines ? Is it consistent throughout the codebase ?
* **Usability**: is your API suitable to be used by many different clients, some with memory constraints, some with perfomance issues ? Can your API be used on different type of devices like web, mobile, or IoT ?
* **Security**: are there any obvious vulnerability ?
* **Testing**: how thorough are the automated tests ? Will they be difficult to change if the requirements of the application were to change ? Are there some unit or some integration tests ? We're not looking for full coverage (given time constraint) but just trying to get a feel for your testing skills.
* **Technical choices**: do choices of libraries, databases, architecture etc. seem appropriate for the chosen application ?
* **Documentation**: is your API documented ? With a decent handling of http status code ? With curl examples ? Does it expose a modern documentation (swagger-like) ?
* **Production-readiness**: does the code include monitoring ? logging ? proper error handling ?

Bonus point (those items are optional):

* **Scalability**: will technical choices scale well ? If not, is there a discussion of those choices in the result.md ?
* **Authentication**: does the API allows user authentication ?

## Notes

* We're evaluating you on completion of the scenario, but also on style. Do you commit often? Did you make code that you would enjoy reading ? Things like that are important to our team 👊
* We expect you to provide a code as you would do in any professional situation, with future maintenability and continuous deployment constraints in mind, with deployment or installation instructions, ...
* This should be fun! If you're stuck on an instruction or if something needs clarification, you can start discussion on the repository board.


## Done ?

Great job! When you're all finished, open a PR and we'll review it together 🙌

## Thanks

* http://onqlave.com
