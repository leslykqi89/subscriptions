# Subscription API

This project integrates Go with MongoDB, to create a subscritpion RestAPI. The enpoints included are:

* POST /subscriptor
    * Creates a new subscriptor in the database
*	GET /subscriptor/:id
    * Obtains the subscriptor based on the id
*	GET /subscriptors
    * Returns all the subscriptors
*	PUT /subscriptor/:id
    * Update the subscriptor with the information provided in the request body.
*	DELETE /subscript/:id
    * Removes the subscriptor from the database.
