# Exoplanet Microservice

# Description

A Golang microservice for managing exoplanets, providing CRUD operations and fuel estimation for space voyages.

## Features

1. **Add an Exoplanet**
2. **List Exoplanets**
3. **Get Exoplanet by ID**
4. **Update Exoplanet**
5. **Delete Exoplanet**
6. **Fuel Estimation**

## Exoplanet Types

- **Gas Giant**: Composed of gaseous compounds.
- **Terrestrial**: Earth-like, rocky, larger than Earth.

## Fuel Estimation Formula
- `d`: distance (light years)
- `g`: gravity
- `c`: crew capacity

Gravity Calculation:
- **Gas Giant**: `g = 0.5 / r^2`
- **Terrestrial**: `g = m / r^2`
- Constraints: `10 < d < 1000`, `0.1 < m < 10`, `0.1 < r < 10`

## Getting Started

### Prerequisites

- Go 1.16+
- Docker

## Setup
1. Clone the repository
2. Build and run the Docker container:
   ```bash
   docker build -t exoplanet-service .
   docker run -p 8080:8001 exoplanet-service
