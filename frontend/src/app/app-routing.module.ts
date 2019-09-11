import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { FiletransfertableComponent } from './components/filetransfertable/filetransfertable.component';
import { BascamptableComponent } from './components/bascamptable/bascamptable.component';

const routes: Routes = [
  {path: '', component: LoginComponent},
  {path: 'filetransfertable/:urno', component: FiletransfertableComponent},
  {path: 'basecamptable', component: BascamptableComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
