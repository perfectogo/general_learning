#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <time.h>

// Define the structure for a Racer
struct Racer {
    char name[50];
    int speed;
    int distance_covered;
};

// Define the structure for a RaceCar
struct RaceCar {
    char driverName[50];
    char raceCarColor[20];
    int totalLapTime;
};

// Define the structure for a Race
struct Race {
    int numberOfLaps;
    int currentLap;
    char firstPlaceDriverName[50];
    char firstPlaceRaceCarColor[20];
};

// Function to initialize racers
void initializeRacers(struct RaceTrack *raceTrack) {
    raceTrack->num_racers = 0;

    // Add racers to the race track
    // You can modify this function to add racers dynamically if you want
    // For now, let's add some example racers
    struct Racer racer1 = {"Speedy", 5, 0};
    struct Racer racer2 = {"Swift", 4, 0};
    struct Racer racer3 = {"Furious", 6, 0};

    raceTrack->racers[raceTrack->num_racers++] = racer1;
    raceTrack->racers[raceTrack->num_racers++] = racer2;
    raceTrack->racers[raceTrack->num_racers++] = racer3;
}

// Function to print race status
void printRaceStatus(struct RaceTrack *raceTrack) {
    printf("Current Race Status:\n");
    for (int i = 0; i < raceTrack->num_racers; i++) {
        printf("%s has covered %d meters.\n", raceTrack->racers[i].name, raceTrack->racers[i].distance_covered);
    }
    printf("----------------------------\n");
}

// Function to print race introduction
void printIntro() {
    printf("Welcome to our main event digital race fans!\n");
    printf("I hope everybody has their snacks because we are about to begin!\n");
}

// Function to print race countdown
void printCountDown() {
    printf("Racers Ready! In...\n");
    for (int i = 5; i >= 1; i--) {
        printf("%d\n", i);
    }
    printf("Race!\n");
}

// Function to print first place after a lap
void printFirstPlaceAfterLap(struct Race race) {
    printf("After lap number %d\n", race.currentLap);
    printf("First Place Is: %s in the %s race car!\n", race.firstPlaceDriverName, race.firstPlaceRaceCarColor);
}

// Function to print congratulatory message for the winner
void printCongratulation(struct Race race) {
    printf("Let's all congratulate %s in the %s race car for an amazing performance.\n", race.firstPlaceDriverName, race.firstPlaceRaceCarColor);
    printf("It truly was a great race and everybody have a goodnight!\n");
}

// Function to calculate time taken by a race car to complete a lap
int calculateTimeToCompleteLap() {
    int speed = rand() % 3 + 1;        // Random value between 1 and 3
    int acceleration = rand() % 3 + 1; // Random value between 1 and 3
    int nerves = rand() % 3 + 1;       // Random value between 1 and 3

    return speed + acceleration + nerves;
}

// Function to update the total lap time of a race car
void updateRaceCar(struct RaceCar *raceCar) {
    raceCar->totalLapTime += calculateTimeToCompleteLap();
}

// Function to update the first place driver and car color
void updateFirstPlace(struct Race *race, struct RaceCar *raceCar1, struct RaceCar *raceCar2) {
    if (raceCar1->totalLapTime <= raceCar2->totalLapTime) {
        strcpy(race->firstPlaceDriverName, raceCar1->driverName);
        strcpy(race->firstPlaceRaceCarColor, raceCar1->raceCarColor);
    } else {
        strcpy(race->firstPlaceDriverName, raceCar2->driverName);
        strcpy(race->firstPlaceRaceCarColor, raceCar2->raceCarColor);
    }
}

// Function to start the race
void startRace(struct RaceCar *raceCar1, struct RaceCar *raceCar2) {
    struct Race race;
    race.numberOfLaps = 5;
    race.currentLap = 1;
    strcpy(race.firstPlaceDriverName, "");
    strcpy(race.firstPlaceRaceCarColor, "");

    for (int lap = 1; lap <= race.numberOfLaps; lap++) {
        updateRaceCar(raceCar1);
        updateRaceCar(raceCar2);
        updateFirstPlace(&race, raceCar1, raceCar2);
        printFirstPlaceAfterLap(race);
        race.currentLap++;
    }

    printCongratulation(race);
}

int main() {
    // Seed the random number generator
    srand(time(NULL));

    // Create and initialize two race cars
    struct RaceCar raceCar1 = {"Cosmo", "orange", 0};
    struct RaceCar raceCar2 = {"Luna", "blue", 0};

    // Print the race introduction
    printIntro();

    // Print the countdown
    printCountDown();

    // Start the race simulation
    startRace(&raceCar1, &raceCar2);

    return 0;
}
