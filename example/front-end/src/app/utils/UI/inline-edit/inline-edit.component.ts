import {Component, Input, Output, EventEmitter, Renderer, ElementRef,ViewChild} from '@angular/core';
//import {ControlValueAccessor, NG_VALUE_ACCESSOR} from "@angular/common";

@Component({
    selector: 'inline-edit',
    styles: [require('./inline-edit.scss')],
    template: require('./inline-edit.html')
})
export class InlineEditComponent {

    @ViewChild('inputEle') inputEle : ElementRef;

    @Input() value: any;
    @Input() type: any;
    @Input() isEdit: boolean;
    @Input() placeHolder: any;
    @Output('onUpdate') model = new EventEmitter();

    private preValue: string = '';
    
    constructor(public renderer: Renderer, public elementRef: ElementRef) { }

    ngOnInit() {
        this.type = (this.type)?this.type:"text";
        this.placeHolder = (this.placeHolder)?this.placeHolder:"";
        this.isEdit = (this.isEdit)?this.isEdit:false;
    }

    isNotEditing() {
        return this.isEdit
    }

    isEditing() {
        return !this.isEdit;
    }
    // Method to display the inline edit form and hide the <a> element
    edit(value) {
        this.preValue = value;  // Store original value in case the form is cancelled
        this.isEdit = true;
        // Automatically focus input
        //this.renderer.invokeElementMethod(this.inlineEditControl.nativeElement, 'focus', []);
        setTimeout(_ => this.renderer.invokeElementMethod(this.inputEle.nativeElement, 'focus', []));

    }

    // Method to display the editable value as text and emit save event to host
    onUpdate(value) {
        this.value = value;
        if(value == "" || value == undefined || value == null){
            this.value = this.preValue;
        }
        this.model.emit(this.value);
        this.isEdit = false;
    }
}