# Components

Currently a little tricky to use.  Here's what you need to do:

- Add the component's model to your main state, e.g.,
```
State: MyAppModel{
	MyComponentState: myComponent.Model{},
},
```

- Parse the component's template, e.g. `existingTemplates.Must(Templates.ParseGlob("../components/*/*.html"))`

- Override the blank message handler for each message, probably in an init() function somewhere.

- Merge in the component's message map, e.g. 
```
	gotea.App.Messages.
		MergeMap(myComponent.Messages)
```

- Use it: 
```
{{ template "mycomponent.html" .MyComponentState }}

```