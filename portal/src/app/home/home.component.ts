import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ExamService } from '../shared/exam.service';
import { NavbarComponent } from '../navbar/navbar.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  constructor(public examService:ExamService, private router: Router) { }

  ngOnInit() {
    localStorage.setItem('qnProgress', "0");
    localStorage.setItem('qns', "");
    localStorage.setItem('seconds', "0");
    var paricipant=this.examService.getParticipant();
    if(paricipant.type !="user"){
      this.examService.SignOut();
    }

    //get category
  }

  LauchTest(exam :string ){
    localStorage.setItem('testLaunchCategory', exam);
    this.router.navigate(['/exam']);
  }
}
