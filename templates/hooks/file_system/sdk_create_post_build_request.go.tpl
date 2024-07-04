    // This is an idempotency token required in the API call...
    //input.SetCreationToken(getIdempotencyToken())

    //For AWS-SDK-GO-V2
    token := getIdempotencyToken()
	input.CreationToken = &token