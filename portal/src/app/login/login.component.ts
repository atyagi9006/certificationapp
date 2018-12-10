import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../shared/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  emailPattern = "^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$";

  constructor(private userService: UserService,private router:Router) { }

  OnSubmit(email:string, password:string){
    this.userService.authenticate(email,password).subscribe(
      (data:any)=>{
        localStorage.clear();
        localStorage.setItem('participant',JSON.stringify(data))
        if(data.type=="admin"){
          this.router.navigate(['/admin']);
          return
        }
        this.router.navigate(['/home']);
      }
    );
  }
  ngOnInit() {
  }
  OnsignUp(){
    this.router.navigate(['/register']);
  }
}
