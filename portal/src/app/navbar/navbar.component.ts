import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ExamService } from '../shared/exam.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  constructor(private router:Router,private examService:ExamService) { }

  ngOnInit() {
  }

  SignOut() {
    this.examService.SignOut();
  }
}
