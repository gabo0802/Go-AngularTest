import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']

  //you can use:
  //changeDetection: ChangeDetectionStrategy.OnPush
  //to change the detection strategy for the setTimeout() call
  
})
export class AppComponent {
  title = 'My Test :)';
}
