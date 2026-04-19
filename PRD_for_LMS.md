# Product Requirements Document (PRD)
## Multi-Tenant Learning Management System (LMS) for Coaching Institutions

---

## Executive Summary

The Indian coaching ecosystem is highly fragmented, operationally manual, and digitally inconsistent. While students increasingly expect seamless digital learning experiences, most coaching institutions lack robust systems to manage academic operations at scale.

This PRD outlines the requirements for a **multi-tenant, institute-centric Learning Management System (LMS)** that enables coaching institutes to digitize and scale their academic workflows while maintaining operational autonomy. The platform is designed to serve three primary stakeholders—**Platform Admins, Institute Admins/Teachers, and Students**—with clearly defined responsibilities, access boundaries, and success metrics.

The MVP focuses on **academic operations**, not monetization, ensuring rapid adoption and strong product-market fit.

---

## 1. Problem Statement

Coaching institutions today face three core challenges:

1. **Operational Fragmentation**  
   Attendance, lectures, notes, doubts, and test materials are managed across disconnected tools.

2. **Lack of Scalable Academic Infrastructure**  
   Most existing solutions are either too generic or too complex for coaching institutes.

3. **Poor Student Experience**  
   Students lack a unified, intuitive platform to access lectures, materials, and performance data.

This LMS aims to address these gaps through a **role-based, institute-isolated, and scalable system**.

---

## 2. Product Vision

> *Build the default academic operating system for coaching institutions—simple for small institutes, powerful for large ones, and delightful for students.*

---

## 3. Design Principles

1. **Institute Autonomy**  
   Each institute operates independently with strict data isolation.

2. **Role Clarity**  
   Every feature maps clearly to a single user role.

3. **Operational Simplicity**  
   Minimize cognitive and operational overhead for admins and teachers.

4. **Student-First Consumption**  
   Content discovery, access, and tracking must be frictionless.

5. **Scalable by Default**  
   Architecture must support high read traffic and rapid growth.

---

## 4. User Personas

### 4.1 Platform Admin
Owns platform integrity, institute onboarding, and high-level governance.

### 4.2 Institute Admin / Teacher
Runs day-to-day academic operations for a coaching institute.

### 4.3 Student
Consumes content, tracks attendance, and interacts academically through doubts.

---

## 5. Functional Requirements

---

## 5.1 Platform Admin Capabilities

### 5.1.1 Institute Lifecycle Management
- Onboard new coaching institutes
- Deactivate or permanently remove institutes
- Each institute must have a **globally unique, immutable `institute_id`**

**Institute Metadata**
- Institute ID
- Name
- Contact information
- Status (Active / Inactive)
- Created timestamp

---

### 5.1.2 Institute Visibility
- View institute profiles
- View aggregated metrics:
  - Total students
  - Total batches
  - Total teachers
- Modify non-academic institute metadata only

---

## 5.2 Institute Admin / Teacher Capabilities

---

### 5.2.1 Student Management
- Add, update, and deactivate students

**Student Profile**
- Full name
- Phone number (primary identifier)
- Email (optional)
- Enrollment status
- Associated batches

---

### 5.2.2 Batch Management
- Create, update, and archive batches

**Batch Attributes**
- Batch ID
- Batch name
- Academic year
- Subjects
- Assigned teachers
- Enrolled students

A student may be enrolled in multiple batches.

---

### 5.2.3 Lecture & Class Management
Each batch contains a sequence of lectures (classes).

For each lecture:
- Define lecture date and time
- Upload one or more videos
- Upload lecture-specific notes
- Lectures must be **chronologically ordered** by class time

---

### 5.2.4 Attendance Management
- Attendance is recorded **per lecture**
- Admin can mark and update attendance before final lock
- Attendance data must be:
  - Stored at lecture level
  - Aggregated at batch and student level

---

### 5.2.5 Study Material Management
Admins can upload batch-level materials including:
- Lecture notes
- Videos
- Previous Year Questions (PYQs)
- Test papers
- Test results
- Custom material types

**Material Constraints**
- Tagged by type
- Visible only to enrolled students
- Searchable and filterable

---

### 5.2.6 Doubt Resolution
- Students can ask **one doubt per lecture**
- Admin/Teacher can provide **one reply**
- No real-time or threaded chat
- Doubts are:
  - Lecture-scoped
  - Visible only to the concerned student and admin

---

## 5.3 Student Capabilities

---

### 5.3.1 Authentication
- Login using:
  - Institute ID
  - Phone number
  - OTP verification
- No password-based login

---

### 5.3.2 Dashboard
- View enrolled batches
- Access upcoming and past lectures
- Navigate batch-wise materials

---

### 5.3.3 Profile Management
- View and update editable personal details

---

### 5.3.4 Lecture Consumption
For each lecture, students can:
- Watch videos
- View/download notes
- Access supplementary materials
- View attendance status
- Add personal notes (private)

---

### 5.3.5 Doubt Asking
- Ask one doubt per lecture
- View teacher/admin response
- No follow-up interactions

---

## 6. Non-Functional Requirements

---

### 6.1 Security
- OTP-based authentication
- Role-based authorization
- Strict institute-level data isolation
- Secure media access via signed URLs or tokens

---

### 6.2 Performance
- Page load time < 2 seconds (excluding video buffering)
- Optimized for read-heavy workloads

---

### 6.3 Scalability
- Designed to support:
  - 10,000+ institutes
  - 1M+ students
  - High concurrent video consumption

---

### 6.4 Availability
- Target uptime: 99.9%
- Graceful degradation for video services

---

## 7. Audit & Data Integrity
- Audit logs for:
  - Student enrollment changes
  - Attendance updates
  - Material uploads and deletions
- Soft deletes preferred over hard deletes

---

## 8. Success Metrics (MVP)

- Institute onboarding completion rate
- Daily active students
- Lecture completion rate
- Attendance accuracy
- Average doubt resolution time

---

## 9. Explicitly Out of Scope (MVP)

- Live classes
- Payments and subscriptions
- Parent access
- Advanced analytics
- Push or SMS notifications

---

## 10. Future Roadmap Considerations

- Live teaching and recordings
- Monetization and subscriptions
- Parent dashboards
- Learning analytics and insights
- Notification infrastructure

---

## Closing Note

This PRD intentionally prioritizes **academic workflow excellence over feature breadth**. The goal is to establish a robust foundation upon which monetization, analytics, and engagement layers can be added without architectural rework.

