# Subscription API

This project integrates Go with MongoDB, to create a subscritpion RestAPI. The enpoints included are:

* POST /subscriber
    * Creates a new subscriber in the database
*	GET /subscriber/:id
    * Obtains the subscriber based on the id
*	GET /subscribers
    * Returns all the subscribers
*	PUT /subscriber/:id
    * Update the subscriber with the information provided in the request body.
*	DELETE /subscript/:id
    * Removes the subscriber from the database.
