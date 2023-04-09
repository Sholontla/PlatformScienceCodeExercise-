import math
import numpy as np


class Shipment:
    def __init__(self, destination):
        self.destination = destination


class Driver:
    def __init__(self, name):
        self.name = name


class Assignment:
    def __init__(self, shipment, driver):
        self.shipment = shipment
        self.driver = driver


class ShipmentAssigner:
    def __init__(self, shipments, drivers):
        self.shipments = [Shipment(s) for s in shipments]
        self.drivers = [Driver(d) for d in drivers]

    def assign(self):
        assignments = []
        used_drivers = set()

        for s in self.shipments:
            best_score = -math.inf
            best_driver = None

            for d in self.drivers:
                if d in used_drivers:
                    continue

                ss = self.calculate_suitability_score(s, d)
                if ss > best_score:
                    best_score = ss
                    best_driver = d

            assignments.append(Assignment(s, best_driver))
            used_drivers.add(best_driver)

        return assignments

    def calculate_suitability_score(self, s, d):
        sn = len(s.destination)
        dn = len(d.name)

        # calculate base SS
        if sn % 2 == 0:
            base_ss = self.count_vowels(d.name) * 1.5
        else:
            base_ss = self.count_consonants(d.name)

        # apply factor bonus
        if self.find_common_factor(sn, dn) > 1:
            base_ss *= 1.5

        return base_ss

    def count_vowels(self, s):
        count = 0
        for r in s:
            if self.is_vowel(r):
                count += 1
        return count

    def count_consonants(self, s):
        count = 0
        for r in s:
            if not self.is_vowel(r):
                count += 1
        return count

    def is_vowel(self, r):
        return r in ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']

    def find_common_factor(self, a, b):
        for i in range(2, min(a, b) + 1):
            if a % i == 0 and b % i == 0:
                return i
        return 1


class PlatformScienceCodeExercise:
    def __init__(self, shipments, drivers):
        self.shipments = np.array(shipments)
        self.drivers = np.array(drivers)

    def run(self):
        # sample data
        assigner = ShipmentAssigner(
            self.shipments.tolist(), self.drivers.tolist())
        assignments = assigner.assign()

        for i, a in enumerate(assignments):
            print(
                f"Day {i+1}: Shipment to {a.shipment.destination} assigned to driver {a.driver.name}")


if __name__ == "__main__":

    shipments = np.array(["Maple Street", "Oak Lane", "Elm Road"])
    drivers = np.array(["John Smith", "Jane Doe", "Mike Johnson"])
    process = PlatformScienceCodeExercise(shipments, drivers)
    process.run()
