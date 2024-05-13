Feature: Capture Pokemon

Background:
  Given a Pokemon appears

  Scenario: Capture Pokemon that is friendly
    Given the Pokemon is friendly
    When the trainer tries to capture the Pokemon
    Then the Pokemon is captured

  Scenario: Capture Pokemon that is not friendly
    Given the Pokemon is not friendly
    When the trainer tries to capture the Pokemon
    Then the Pokemon escapes

  Scenario: The trainer has no pokeballs
    Given the trainer has no pokeballs
    When the trainer tries to capture the Pokemon
    Then the Pokemon escapes
