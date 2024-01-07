Try to design a basic schema for Netflix.
================================================================
**UseCases**:
================================================================
- Every user has an email and a password.
- Every User can have up to 5 profiles.
- Every profile has a name and a type.
- Netflix has videos.
- Every video has a name, description, and list of actors.
- Each Actor has a name and list of videos he/she is a part of.
- For every video, for every user we want to store the status i.e. Currently Watching or Watched.
- If currently watching, store the last currently watched timestamp.