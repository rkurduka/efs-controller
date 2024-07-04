    // This is an idempotency token required in the API call...
    //input.SetClientToken(getIdempotencyToken())  

    //For AWS-SDK-GO-V2
    token := getIdempotencyToken()
	input.ClientToken = &token

