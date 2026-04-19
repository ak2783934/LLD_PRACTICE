* Design database for Twitter with given requirements:
    * User should be able to - follow unfollow
    * User should be able to post tweets
    * User should be able to see his profile page - his tweets only
    * User should be able to see Home Screen, everyones tweets..
    * Add Like feature
    * we should be able to see the total number of likes, when a user clicks on it, he should be able to see who all have liked his tweets
    * After giving ER diagram for above requirements, additional requirement was added.
    * How to handle celebrity problem, how to effectively find number of likes on celebrity posts: for later....



Tables/Entities -> think from the POV of proper database. 
- also indexing which might be required. 


index on id

User{
    id string
    name string
    photo string
}

// when unfollow, we delete this entry. IMP. 
index on userID 
Follow{
    userID string
    followeeID string
    createdAt string
}

index on userID
Tweet{
    id string
    userID string 
    content []byte
}

Index on tweetID
Like{
    id string
    tweetID string
    likedBy string
}