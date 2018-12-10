import { Component, OnInit } from '@angular/core';
import { UserService } from '../shared/user.service';
import { User } from '../shared/user.model';
import { ExamAttempt } from '../shared/ExamAttempt';


@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {
  users: User[];
  examAttemptList: ExamAttempt[];
  candidates: any[];
  constructor(public userService: UserService) {
  }

  ngOnInit() {
    var loggedInUser = this.userService.getLoggedInUser();
    if (loggedInUser.type != "admin") {
      this.userService.Signout();
    }
    this.userService.getUsers().subscribe(
      (data: any) => {
        //console.log(data)
        this.users = data;
       /*  for (let user of data) {
          this.candidates[user.userId] = this.getCandidateDetails(user.userId);
        } */
      }
    );
  }

  getCandidateDetails(userId: string) {
    console.log("requested by.." + userId);
    var result;
    this.userService.getCandidate(userId).subscribe(
      (data: any) => {
        console.log(data);
        result = data;
      }
    );
    return result;
  }
  showCandidateDetails(userId: string) {
    console.log("show requested by.." + userId);
    this.examAttemptList=this.getCandidateDetails(userId).examAttemptList;
  }
}
