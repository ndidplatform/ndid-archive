Feature: Authentication
  As a relying party
  I want to authenticate the user's identity

  @mock_rp
  Scenario: Getting a request ID
    When I call the send request API
    Then I should get a request ID

  @mock_rp
  Scenario: Calling the send_request_to_id API twice with same reference ID
    When I call the send request API with reference ID "e3cb44c9-8848-4dec-98c8-8083f373b1f7"
    And I call the send request API with reference ID "e3cb44c9-8848-4dec-98c8-8083f373b1f7"
    Then received request IDs should be the same

  @mock_rp
  Scenario: Calling the send_request_to_id API twice with different reference ID
    When I call the send request API with reference ID "ee2a741b-056c-446f-bf4c-a73c0af29ed0"
    And I call the send request API with reference ID "ddcaf9ae-a6ac-4bf9-ba9c-ce2e75601077"
    Then received request IDs should be different
