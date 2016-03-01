import 'rxjs/Rx';
import {Component} from 'angular2/core';
import {NgForm}    from 'angular2/common';
import {Injectable, Injector} from 'angular2/core';
import {HTTP_PROVIDERS, Http} from 'angular2/http';


@Component({
    // Declare the tag name in index.html to where the component attaches
    selector: 'autotune-curl',
    viewProviders: [HTTP_PROVIDERS],
    // Location of the template for this component
    template: `
<div>
<label>Server Profile:</label>
<select [(ngModel)]="selectedProfile" class="form-control">
<option *ngFor="#p of profiles" [value]="p">{{p}}</option>
</select>
</div>
<br>
<code>\\curl -sSL https://acksin.com/autotune/install.sh | bash -s {{selectedProfile}}</code>`
})
export class AutotuneCurl {
    constructor(http: Http) {
        http.get('js/profiles.json').map(res => res.json())
            .subscribe(profiles => this.profiles = profiles);
    }

    // Default profile to use.
    selectedProfile: string = 'apache';
}
