# NWG DE Technical Assessment

![go](https://badges.aleen42.com/src/golang.svg) ![Build](https://github.com/JOHNEPPILLAR/nwg-de/actions/workflows/go.yml/badge.svg) [![GitHub tag](https://img.shields.io/github/tag/johneppillar/nwg-de?include_prereleases=&sort=semver&color=blue)](https://github.com/johneppillar/nwg-de/releases/)
[![License](https://img.shields.io/badge/License-MIT-blue)](#license)

# Scenario
During this exercise, your customer is a car rental firm owner. The car rental firm owner doesn’t know anything about technology; however, they do know what they need and they are working on a tight deadline. They give you the requirements 3 days before meeting you to discuss your solution (i.e. Your Technical Interview). You will need to send them your solution 1 day before the meeting to be reviewed by their Engineers. (i.e. The DE Panel) They have a set of requirements in the form of stories which you will implement one at a time. You are given 2 stories to complete and you are expected to spend no more than 1-2 hours for the exercise. You will need to send your solution to the customer once it is ready no later than 1 day before your meeting. Please complete the stories the best you can and prepare to discuss your solution and choices with your interviewer. They will assess your knowledge and thought process behind the implementation decisions. It is not important to implement all corner cases, however it is important to show there are things you might have considered if this was "production" code, and you can talk about those with your assessors.

# About this service

This service is written in GoLang v 1.20 with the software architecture and design principals being:
- Hexagonal architecture pattern
- 12 factor app patterns
- Infra as code
- GitOps
- Design for HA

This scenario is a classic technical interview question related to OO programming. I've not chosen a classic OO design as its rather outdated in my personal view, I've selected Hexagonal Architecture as the design approach as it encompasses the core OO principals and extends them further with domain driven design and layer abstraction. If this is new to you I suggest reading [Hexagonal Architecture, there are always two sides to every story](https://medium.com/ssense-tech/hexagonal-architecture-there-are-always-two-sides-to-every-story-bc0780ed7d9c)

Operational consideration:
- All std out logs feed into grafana loki
- Traefik used for kube ingress
- Vault for secrets management (using vault sidecar as service is in a kube pod)
- Flux for GitOps
- Github actions used for CI/CD
- Docker container(s) run in a kubernetes cluster
- Set app to be HA (3 node deployment)

No UI device was specified thus this is only the backend api service, plus I ran out of time! 

The data repository used is mongoDB. The driving https Adaptor is Gin which is great for high volume requests. Middleware exists to rate limit, add secure headers and also api's sit behind a bearer token for auth 😉

Health check end points:
- No auth'ed health check end point: /v1/health 
- Auth'ed Prometheus metrics end point: /v1/metrics

API tests are located in the postman folder (https://github.com/JOHNEPPILLAR/nwg-de/postman)

## User Story 1 - Finding a car to rent

As a car rental company, I want to match a potential renter to a car that I have to rent...
...So that I can give a list of cars to a potential client.

- Acceptance criteria: one method returning a list of matching cars
- Design criteria: the method should be thread safe (allowing it to be called from multiple threads)

API: (GET) /v1/vehicle/available
Comments: Client side filtering for vehicle options is faster than calling a filtering api
Added extras: Add vehicle end point (POST) /v1/vehicle

## Story 2 - Finding an available car to be rented

As a car renter I want to know if a car is available to be rented on the dates I need...
...So that I can provide a list of cars that are available on the dates given by the renter

Acceptance criteria: the matching criteria for a renter to rent a car should include a from date and to date
Acceptance criteria: the car renter should not be shown any cars that are booked in the period that is supplied
Acceptance criteria: one method returning a list of matching cars (with the filter having changed)
Design criteria: the method should be thread safe (allowing it to be called from multiple threads)

API: (POST) /v1/vehicle/available

## Story 3 - Booking a car

As a car renter I want to book a car which has been shown to me as being available...
...So that I can have a car available to me to use during the rental period

- Acceptance criteria: the car rental should be stored in an object model
- Acceptance criteria: there should not be able to have overlapping car rentals for the same car
- Acceptance criteria: two renters should not be able to book the same car at the same time for an overlapping period
- Acceptance criteria: one method allowing the car renter to book a car for a period
- Design criteria: the in memory storage should consider thread safety, all access to and from it should therefore be thread safe

API: (POST) /v1/booking
Added extras: Cost for hire is calculated at time of booking