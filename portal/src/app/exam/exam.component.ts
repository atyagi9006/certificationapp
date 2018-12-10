import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ExamService } from '../shared/exam.service';

@Component({
  selector: 'app-exam',
  templateUrl: './exam.component.html',
  styleUrls: ['./exam.component.css']
})
export class ExamComponent implements OnInit {

  constructor(private router: Router, public examService: ExamService) { }

  ngOnInit() {
    var paricipant=this.examService.getParticipant();
    if(paricipant.type !="user"){
      this.examService.SignOut();
    }
    if (parseInt(localStorage.getItem('seconds')) > 0) {
      this.examService.seconds = parseInt(localStorage.getItem('seconds'));
      this.examService.qnProgress = parseInt(localStorage.getItem('qnProgress'));
      this.examService.qns = JSON.parse(localStorage.getItem('qns'));
      if (this.examService.qnProgress == 10)
        this.router.navigate(['/result']);
      else
        this.startTimer();
    }
    else {
      this.examService.seconds = 0;
      this.examService.qnProgress = 0;
      this.examService.getQuestions().subscribe(
        (data: any) => {
          console.log(data)
          this.examService.qns = data;
          this.startTimer();
        }
      );
    }
  }
  startTimer() {
    this.examService.timer = setInterval(() => {
      this.examService.seconds++;
      localStorage.setItem('seconds', this.examService.seconds.toString());
    }, 1000);
  }


  Answer(qID, choice) {
    this.examService.qns[this.examService.qnProgress].answer = choice;//added one more property
    localStorage.setItem('qns', JSON.stringify(this.examService.qns));
    this.examService.qnProgress++;
    localStorage.setItem('qnProgress', this.examService.qnProgress.toString());
    if (this.examService.qnProgress == 10) {
      clearInterval(this.examService.timer);
      this.router.navigate(['/result']);
    }
  }

}
