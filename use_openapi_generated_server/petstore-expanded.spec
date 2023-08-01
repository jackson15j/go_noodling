Feature: Petstore Extended API
  Background:
    Given openapi ./petstore-expanded-v3.yaml

  Scenario Outline: Create Pet
    When POST /pets
    Then status 201
    Examples:
      | name        | tag |
      | Molly       | dog |
      | Mr. Tinkles | cat |

  Scenario Outline: Get Pet Success
    When GET /pets/(id:number)
    Then status 200
    Examples:
      |   id | name | tag |
      | 1000 | Rex  | dog |

  Scenario Outline: Get Pet Not Found Error
    When GET /pets/5
    Then status 404

  Scenario Outline: Delete Pet Success
    When DELETE /pets/(id:number)
    Then status 204
    Examples:
      |   id | name | tag |
      | 1000 | Rex  | dog |

  Scenario Outline: Delete Pet Not Found Error
    When DELETE /pets/5
    Then status 404

