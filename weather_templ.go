// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

func Base() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:#f0f0f0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-family:Arial, sans-serif;`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:#333;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:100vh;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-sizing:border-box;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Base`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func GetLocation(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_GetLocation_1bc0`,
		Function: `function __templ_GetLocation_1bc0(id){if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition((position) => {
            const latitude = position.coords.latitude;
            const longitude = position.coords.longitude;
            console.log(latitude, longitude);
            const element = document.getElementById(id)
            element.setAttribute("data", ` + "`" + `${latitude},${longitude}` + "`" + `)
            return
        },
        function(error) {
            console.log("Unable to retrieve your location", error)
        }
        );
    }else {
        console.log("Geolocation is not supported by this browser.");
    }
}`,
		Call:       templ.SafeScript(`__templ_GetLocation_1bc0`, id),
		CallInline: templ.SafeScriptInline(`__templ_GetLocation_1bc0`, id),
	}
}

func FetchWeather(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_FetchWeather_8aa0`,
		Function: `function __templ_FetchWeather_8aa0(id){const element = document.getElementById(id)
    const dataValue = element.getAttribute("data");
    console.log("dataValue:  ",dataValue);
    const [latitude, longitude] = dataValue.split(",");

    const apikey = "d3f6e48d659248fdb9464301241002"
    const url = ` + "`" + `https://api.weatherapi.com/v1/current.json?key=${apikey}&q=Lalitpur&aqi=no` + "`" + `;

    fetch(url)
        .then(response => response.json())
        .then(data => {
            console.log('Weather Data: ',data);
            textContent = ` + "`" + `${data.location.country} <br></br>
            ${data.current.condition.text} <br></br>
            ${data.current.feelslike_c}°C <br></br>
            <img src="${data.current.condition.icon}"/>` + "`" + `
            element.innerHTML = textContent
        })   
        .catch(error => {
            console.log("Error fetching weather data", error);
        });

}`,
		Call:       templ.SafeScript(`__templ_FetchWeather_8aa0`, id),
		CallInline: templ.SafeScriptInline(`__templ_FetchWeather_8aa0`, id),
	}
}

func WeatherForcast(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><title>Weather</title></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 = []any{Base()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, GetLocation("data-div"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body onload=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.ComponentScript = GetLocation("data-div")
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, FetchWeather("data-div"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 templ.ComponentScript = FetchWeather("data-div")
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Hello, ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `weather.templ`, Line: 66, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><div id=\"data-div\"></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
