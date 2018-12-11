import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from './user.model';
import { Candidate } from './candidate.model';
import { ExamService } from './exam.service';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  readonly rootURL = 'http://localhost:8080';
  users: User[];
  meassage = "";
  candidate: Candidate;
  constructor(private http: HttpClient,private examService:ExamService) { }

  insertParticipant(name: string, email: string, password: string) {
    var body = {
      "name": name,
      "email": email,
      "password":password
    }
    return this.http.post(this.rootURL + '/signup', body);
  }


  authenticate(email: string, password: string) {
    var req = {
      "email": email,
      "password": password
    }
    return this.http.post('http://localhost:8080/signin', req);
  }
  getUsers() {
    var participant = JSON.parse(localStorage.getItem('participant'));
    var body = {
      "userId": participant.userId
    }
    return this.http.post(this.rootURL + '/users', body);

  }
  getCandidate(userId: string) {
    var body = {
      "candidateId":"candidate_"+userId
    }
    return this.http.post(this.rootURL + '/candidate', body);
    //.toPromise().then(res=> this.candidate = res as Candidate );

  }
  checkEmail(email: string){
    var  body={
      "email": email
    }
    return this.http.post(this.rootURL + '/checkemail', body);

  }
  Signout(){
      this.examService.SignOut();
  }
  getLoggedInUser(){
    return this.examService.getParticipant();
  }

}
