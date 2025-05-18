# Stack Overflow
## Requirements
- Users can post questions, answer questions, and comment on questions and answers.
- Users can vote on questions and answers.
- Questions should have tags associated with them.
- Users can search for questions based on keywords, tags, or user profiles.
- The system should assign reputation score to users based on their activity and the quality of their contributions.
- The system should handle concurrent access and ensure data consistency.



// classes and funcationalities. 

comments and answers should be same entity. 
One original question class should be there. 
If there is edit or any other access pattern, then use locks. Access/modify things. 
Tags and keywords should be same entity. 

Votings? on questions and users? 
How does a user is rated? 
- based on number of upvotes in their answers and questions. 
- Number of question asked? 
- Number of answers given? 
- 


Users {
    id string 

}
CreateUsers
UpdateUsers
Login/Logout X 
UpdateVotes

Query{
}
AskQuestion
EditQuestion
GiveAnswer
EditAnswers
VoteQuestionsAndAnswers -> these votes should also be counted for those users who are the owners
AddTags(open to all function, people based on their views, can add tags)

SearchQuestions(based on tags, title)
SearchQuestionsFromAUser
SearchAnswersFromAUser

Comments is important. 
It can be an array in any of the question or answers. 
They can't be voted for sure. 
Only question or answer can be voted. 


because answers ke andar bhi comments rakhna hai kya?? Like that of in reddit? 
Or just one array of answers and sorted based on time? 
make it simple. 
questions and its answers only. No reply to answers are allowed. May be we can use some sort of field to decide that. 



