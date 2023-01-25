import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { last, lastValueFrom } from 'rxjs';
import { OnInit } from '@angular/core';
interface INewsfeedItem {
  title: string;
  post: string;
}
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']

  //you can use:
  //changeDetection: ChangeDetectionStrategy.OnPush
  //to change the detection strategy for the setTimeout() call
  
})
export class AppComponent implements OnInit {

  public title = ''
  public post = ''

  public newsFeedItems: INewsfeedItem[] = []

    //Creates a singleton instance of the HttpClient
    constructor(
      private httpClient: HttpClient
      ) {}

    async ngOnInit() {
      await this.loadNewsItems()
    }
    //Tutorial uses .toPromise() instead of await lastValueFrom() but the one I used is more modern
    async loadNewsItems() {
      this.newsFeedItems = await lastValueFrom(this.httpClient.get<INewsfeedItem[]>('api/newsfeed'))
    }
    async addPost() {
      await lastValueFrom(this.httpClient.post('api/newsfeed', {
        title: this.title,
        post: this.post}))
      await this.loadNewsItems()
      this.title = ''
      this.post = ''
  }
}
