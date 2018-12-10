import { Component, OnInit } from '@angular/core';
import { ExamService } from '../shared/exam.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-result',
  templateUrl: './result.component.html',
  styleUrls: ['./result.component.css']
})
export class ResultComponent implements OnInit {

  constructor(public examService: ExamService, private router: Router) { }

  ngOnInit() {
    var paricipant=this.examService.getParticipant();
    if(paricipant.type !="user"){
      this.examService.SignOut();
    }
    if (parseInt(localStorage.getItem('qnProgress')) == 10) {
      this.examService.seconds = parseInt(localStorage.getItem('seconds'));
      this.examService.qnProgress = parseInt(localStorage.getItem('qnProgress'));
      this.examService.qns = JSON.parse(localStorage.getItem('qns'));
      this.examService.getAnswers().subscribe(
        (data: any) => {
          this.examService.correctAnswerCount = 0;
          this.examService.qns.forEach((e, i) => {
            if (e.answer == data[e.questionId])
              this.examService.correctAnswerCount++;
            e.correct = data[e.questionId]
          })
        }
      );
     
    }else{
      this.router.navigate(['/exam']);
    }
  }
  OnSubmit() {
    this.examService.submitScore().subscribe(() => {
      this.router.navigate(['/home']);
    });
  }
  restart() {
    localStorage.setItem('qnProgress', "0");
    localStorage.setItem('qns', "");
    localStorage.setItem('seconds', "0");
    this.router.navigate(['/exam']);
  }

}
