import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RegisterComponent } from './register/register.component';
import { ExamComponent } from './exam/exam.component';
import { ResultComponent } from './result/result.component';
import { AuthGuard } from './auth/auth.guard';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { AdminComponent } from './admin/admin.component';

const routes: Routes = [
  { path: 'register', component: RegisterComponent },
  { path: 'exam', component: ExamComponent, canActivate: [AuthGuard] },
  { path: 'result', component: ResultComponent, canActivate: [AuthGuard] },
  { path: 'home', component: HomeComponent,canActivate: [AuthGuard] },
  {path : 'admin', component : AdminComponent,canActivate: [AuthGuard]},
  { path: 'login', component: LoginComponent},
  /* { path: 'signup', component: UserComponent,
    children :[{path:'',component: SignUpComponent}]
  },
  { path: 'login', component: UserComponent,
    children :[{path:'',component: SignInComponent}] 
  }, */
  //{ path: '', redirectTo: '/login', pathMatch: 'full' }
  
  { path: '', redirectTo: '/register', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
