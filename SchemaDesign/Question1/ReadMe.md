Requirements for schema design
================================================================

- A student can have name, email and college.
- Student can login to academy.
- Academy can have multiple courses available.
- A student can enroll in any number of courses.
- Each course has name.
- Each course has course end exams.
- Every exam has a duration.
- One exam can be a part of multiple course.
- Every exam has a specific date for specific course.
- If a student is enrolled in multiple courses, they have to give the exam that many number of times.
- Student can attempt exam only once per course.
- Marks should be stored for every attempt.

================================================================
Step1: Find Entities
================================================================
Student, Course, Exam

Student{
    id
    name
    email
    college
    username
    password
    List<courses>
}

Course{
    id
    name
    List<exams>
}

Exams{
    id
    name
    duration
}

- A student can enroll in any number of courses.
This is relationship  between students and course
StudentCourse{
    studentid
    courseid
}

- Every exam has a specific date for specific course.
This is relationship between Course and Exam. Relationship can also store attributes.
Thus exam date is an attribute of CourseExam Relationship.

CourseExam{
    courseid
    examid
    date
}

- Student can attempt exam only once per course.
This is relation among student and examcourse

Student_CourseExam{
    examcourse
    studentid
    marks
}

================================================================
Step2: Draw Class Diagram
Only those entities will be considered in class which has some attributes
================================================================

1)
Student{
    id
    name
    email
    password
    college
    List<Course>
}

2) 
Course{
    id
    name
    List<Exam>
}

3)
Exam{
    id
    name
    duration
}

4)
CourseExam{
    id
    Course
    Exam
    date
}

5) 
Student_CourseExam{
    id
    Student
    CourseExam
    marks
}

- We will ignore StudentCourse relationship to consider for class diagram, as this relationship does not have any attributes to store.

================================================================
Step3: Relationship Among Entities
================================================================

1) Student and Course
- 1 student can opt for Many course
- 1 course can be opt by Many number of students
- M:M cardinality
- Separate mapping table for student and course

2) Course and Exam
- 1 course can have Many exams
- 1 exam can be a part of Many courses
- M:M cardinality
- Separate mapping table for exam and course

3) CourseExam and Course
- 1 CE can belong to 1 Course
- 1 Course can have Many CE
- M:1 cardinality

4) CourseExam and Exam
- 1 CE can have 1 exam
- 1 Exam can belong to Many CE
- M:1 cardinaltiiy

5) Student_CourseExam and Student
- 1 SEC can belong to 1 student
- 1 student can have Many SEC
- M:1 cardinality

6) Student_CourseExam and CourseExam
- 1 SEC can have 1 CE
- 1 CE can have 1 SEC
- 1:1 cardinality 

================================================================
Step4: Design Schema based on Class Diagram and Cardinality Relationship
================================================================

1) 
Table: Students
Columns: id, name, email, pwd, username, college

2)
Table: StudentsCourseMapping
Columns: id, studentId, courseId

3) 
Table: Course
Columns: id, coursename

4)
Talename: Exams
Columns: id, examname, duration

5)
Table: CourseExamMapping
Columns: id, examid, courseId, date

6)
Table: StudentExamCourseMapping
Columns: id, studentId, examid, courseId, marks




