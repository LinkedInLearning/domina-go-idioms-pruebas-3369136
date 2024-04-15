Feature: Capture Pokémon

Background:
  Given a Pokémon appears

  Scenario: Capture Pokémon that is friendly
    Given the Pokémon is friendly
    When the trainer tries to capture the Pokémon
    Then the Pokémon is captured

  Scenario: Capture Pokémon that is not friendly
    Given the Pokémon is not friendly
    When the trainer tries to capture the Pokémon
    Then the Pokémon escapes

  Scenario: The trainer has no pokeballs
    Given the trainer has no pokeballs
    When the trainer tries to capture the Pokémon
    Then the Pokémon escapes
