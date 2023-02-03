package main

import (
    "github.com/gotk3/gotk3/gtk"
    "log"
    "time"
    "runtime"
)

func main() {
    runtime.LockOSThread()

    gtk.Init(nil)

    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

    if err != nil {
        log.Panic(err)
    }

    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    verticalBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)

    if err != nil {
        log.Panic(err)
    }

    win.Add(verticalBox)


    box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 50)

    if err != nil {
        log.Panic(err)
    }

    verticalBox.Add(box)

    scrolledWindow1, err := gtk.ScrolledWindowNew(nil, nil)

    if err != nil {
        log.Panic(err)
    }

    scrolledWindow1.SetMinContentHeight(250)
    scrolledWindow1.SetMinContentWidth(250)
    scrolledWindow1.SetMaxContentHeight(250)
    scrolledWindow1.SetMaxContentWidth(250)

    box.Add(scrolledWindow1)

    label, err := gtk.LabelNew("foo")

    if err != nil {
        log.Panic(err)
    }

    label.SetSizeRequest(100, 100)
    label.SetMarginStart(10)
    label.SetMarginTop(10)

    scrolledWindow1.Add(label)


    scrolledWindow2, err := gtk.ScrolledWindowNew(nil, nil)

    if err != nil {
        log.Panic(err)
    }

    scrolledWindow2.SetMinContentHeight(250)
    scrolledWindow2.SetMinContentWidth(250)
    scrolledWindow2.SetMaxContentHeight(250)
    scrolledWindow2.SetMaxContentWidth(250)

    box.Add(scrolledWindow2)

    input, err := gtk.TextViewNew()

    if err != nil {
        log.Panic(err)
    }

    input.SetSizeRequest(100, 100)
    input.SetMarginTop(10)
    input.SetMarginEnd(10)

    scrolledWindow2.Add(input)


    errorLabel, err := gtk.LabelNew("")

    if err != nil {
        log.Panic(err)
    }

    verticalBox.Add(errorLabel)


    errorLabelProvider, err := gtk.CssProviderNew()

    if err != nil {
        log.Panic(err)
    }

    errorLabelProvider.LoadFromData("label { color: red; }")

    errorLabelStyleContext, err := errorLabel.GetStyleContext()

    if err != nil {
        log.Panic(err)
    }

    errorLabelStyleContext.AddProvider(errorLabelProvider, gtk.STYLE_PROVIDER_PRIORITY_USER)


    buf, err := input.GetBuffer()

    if err != nil {
        log.Panic(err)
    }

    buf.SetText(`
label {

}
`)

    provider, err := gtk.CssProviderNew()

    if err != nil {
        log.Panic(err)
    }

    timer := time.NewTimer(time.Second)

    buf.Connect("changed", func() {
        timer.Reset(time.Second)
    })

    go func() {
        for {
            select {
                case <-timer.C:
            }

            startIter := buf.GetStartIter()
            endIter := buf.GetEndIter()

            text, err := buf.GetText(startIter, endIter, false)

            if err != nil {
                log.Panic(err)
            }

            styleContext, err := label.GetStyleContext()

            if err != nil {
                log.Panic(err)
            }

            err = provider.LoadFromData(text)

            if err != nil {
                errorLabel.SetLabel(err.Error())
            } else {
                errorLabel.SetText("")

                styleContext.RemoveProvider(provider)
                styleContext.AddProvider(provider, gtk.STYLE_PROVIDER_PRIORITY_USER)
            }
        }
    }()

    win.ShowAll()

    gtk.Main()
}
