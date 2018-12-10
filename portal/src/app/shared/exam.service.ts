import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class ExamService {
  //------------properties--------
  readonly rootURL = 'http://localhost:8080';
  qns: any[];
  seconds: number;
  timer;
  qnProgress: number;
  correctAnswerCount: number = 0;

  //=====helper method=============
  constructor(private http: HttpClient,private router :Router) { }
  displayTimeElapsed() {
    return Math.floor(this.seconds / 3600) + ':' + Math.floor(this.seconds / 60) + ':' + Math.floor(this.seconds % 60);
  }

  //============https methods==============
  
  getQuestions() {
    var  testLaunchCategory= localStorage.getItem('testLaunchCategory');
    var body = {
      "category":testLaunchCategory
    }
    return this.http.post(this.rootURL + '/testlaunch',body);

  }

  getAnswers() {
    var body = this.qns.map(x => x.questionId);
    return this.http.post(this.rootURL + '/answers', body);
  }

  getParticipantName() {
    var participant = JSON.parse(localStorage.getItem('participant'));
    console.log(participant);
    var name=participant.name;
    return name.charAt(0).toUpperCase() + name.slice(1); ;
  }
  getParticipant() {
    var participant = JSON.parse(localStorage.getItem('participant'));
    console.log(participant);
    return participant;
  }
  submitScore() {
    var participant = JSON.parse(localStorage.getItem('participant'));
    var attempt = JSON.stringify(this.qns);
    var testLaunchCategory=localStorage.getItem('testLaunchCategory');
    var body = {
      "candidateID":"candidate_"+participant.userId,
      "examAttemptList": [{
        "categoryId":testLaunchCategory,
        "questionsAttempted": JSON.parse(attempt),
        "score": JSON.stringify(this.correctAnswerCount),
        "timeSpent": JSON.stringify(this.seconds)
      }]
    }
    console.log("the result :  " + body)
    return this.http.post(this.rootURL + "/updatecandidate", body);
  }
  SignOut() {
    localStorage.clear();
    clearInterval(this.timer);
    this.router.navigate(['/login']);
  }
}
