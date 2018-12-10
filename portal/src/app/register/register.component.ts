import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../shared/user.service';
import { Message } from '@angular/compiler/src/i18n/i18n_ast';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  errorMessage=""
  emailPattern = "^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$";
  
  constructor(private userService: UserService,private router:Router,private toastr:ToastrService) { }

  ngOnInit() {
  }
  OnSubmit(name:string , email:string, password :string){

    if (this.checkEmail(email)){
      this.userService.insertParticipant(name,email,password).subscribe(
        (data:any)=>{
          //we can use token based auth here
          /* localStorage.clear();
          localStorage.setItem('participant',JSON.stringify(data)) */
          //this.router.navigate(['/exam']);
         
          this.router.navigate(['/login']);
        }
     );
    }
    
  }
  OnsignIn(){
    this.router.navigate(['/login']);
  }

  search($event) {
    
  }
  checkEmail(email :string){
    var result;
    this.userService.checkEmail(email).subscribe((data:any)=>{
      console.log(data);
      var valid=JSON.parse(data.valid);
      console.log("out"+data.valid)
      if (valid){
        console.log("inside"+data.valid)
       
      }else{
        this.toastr.error("E-mail already in use.");
      } 
      result=valid;
    });
    return result;
  }
}
