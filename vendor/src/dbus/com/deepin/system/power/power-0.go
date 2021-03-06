/*This file is automatically generated by pkg.deepin.io/dbus-generator. Don't edit it*/
package power

import "pkg.deepin.io/lib/dbus"
import "pkg.deepin.io/lib/dbus/property"
import "reflect"
import "sync"
import "runtime"
import "fmt"
import "errors"

/*prevent compile error*/
var _ = fmt.Println
var _ = runtime.SetFinalizer
var _ = sync.NewCond
var _ = reflect.TypeOf
var _ = property.BaseObserver{}

type Power struct {
	Path     dbus.ObjectPath
	DestName string
	core     *dbus.Object

	signals       map[<-chan *dbus.Signal]struct{}
	signalsLocker sync.Mutex

	OnBattery          *dbusPropertyPowerOnBattery
	HasBattery         *dbusPropertyPowerHasBattery
	BatteryPercentage  *dbusPropertyPowerBatteryPercentage
	BatteryStatus      *dbusPropertyPowerBatteryStatus
	BatteryTimeToEmpty *dbusPropertyPowerBatteryTimeToEmpty
	BatteryTimeToFull  *dbusPropertyPowerBatteryTimeToFull
	HasLidSwitch       *dbusPropertyPowerHasLidSwitch
}

func (obj *Power) _createSignalChan() <-chan *dbus.Signal {
	obj.signalsLocker.Lock()
	ch := getBus().Signal()
	obj.signals[ch] = struct{}{}
	obj.signalsLocker.Unlock()
	return ch
}
func (obj *Power) _deleteSignalChan(ch <-chan *dbus.Signal) {
	obj.signalsLocker.Lock()
	delete(obj.signals, ch)
	getBus().DetachSignal(ch)
	obj.signalsLocker.Unlock()
}
func DestroyPower(obj *Power) {
	obj.signalsLocker.Lock()
	defer obj.signalsLocker.Unlock()
	if obj.signals == nil {
		return
	}
	for ch, _ := range obj.signals {
		getBus().DetachSignal(ch)
	}
	obj.signals = nil

	runtime.SetFinalizer(obj, nil)

	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.DBus.Properties',sender='" + obj.DestName + "',member='PropertiesChanged'")
	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='PropertiesChanged'")

	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='BatteryDisplayUpdate'")

	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='BatteryAdded'")

	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='BatteryRemoved'")

	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='LidClosed'")

	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='LidOpened'")

	obj.OnBattery.Reset()
	obj.HasBattery.Reset()
	obj.BatteryPercentage.Reset()
	obj.BatteryStatus.Reset()
	obj.BatteryTimeToEmpty.Reset()
	obj.BatteryTimeToFull.Reset()
	obj.HasLidSwitch.Reset()
}

func (obj *Power) Debug(arg0 string) (_err error) {
	_err = obj.core.Call("com.deepin.system.Power.Debug", 0, arg0).Store()
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj *Power) GetBatteries() (arg0 []dbus.ObjectPath, _err error) {
	_err = obj.core.Call("com.deepin.system.Power.GetBatteries", 0).Store(&arg0)
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj *Power) RefreshBatteries() (_err error) {
	_err = obj.core.Call("com.deepin.system.Power.RefreshBatteries", 0).Store()
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj *Power) ConnectBatteryDisplayUpdate(callback func(arg0 int64)) func() {
	sigChan := obj._createSignalChan()
	go func() {
		for v := range sigChan {
			if v.Path != obj.Path || v.Name != "com.deepin.system.Power.BatteryDisplayUpdate" || 1 != len(v.Body) {
				continue
			}
			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*int64)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(int64))
		}
	}()
	return func() {
		obj._deleteSignalChan(sigChan)
	}
}

func (obj *Power) ConnectBatteryAdded(callback func(arg0 string)) func() {
	sigChan := obj._createSignalChan()
	go func() {
		for v := range sigChan {
			if v.Path != obj.Path || v.Name != "com.deepin.system.Power.BatteryAdded" || 1 != len(v.Body) {
				continue
			}
			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*string)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(string))
		}
	}()
	return func() {
		obj._deleteSignalChan(sigChan)
	}
}

func (obj *Power) ConnectBatteryRemoved(callback func(arg0 string)) func() {
	sigChan := obj._createSignalChan()
	go func() {
		for v := range sigChan {
			if v.Path != obj.Path || v.Name != "com.deepin.system.Power.BatteryRemoved" || 1 != len(v.Body) {
				continue
			}
			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*string)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(string))
		}
	}()
	return func() {
		obj._deleteSignalChan(sigChan)
	}
}

func (obj *Power) ConnectLidClosed(callback func()) func() {
	sigChan := obj._createSignalChan()
	go func() {
		for v := range sigChan {
			if v.Path != obj.Path || v.Name != "com.deepin.system.Power.LidClosed" || 0 != len(v.Body) {
				continue
			}

			callback()
		}
	}()
	return func() {
		obj._deleteSignalChan(sigChan)
	}
}

func (obj *Power) ConnectLidOpened(callback func()) func() {
	sigChan := obj._createSignalChan()
	go func() {
		for v := range sigChan {
			if v.Path != obj.Path || v.Name != "com.deepin.system.Power.LidOpened" || 0 != len(v.Body) {
				continue
			}

			callback()
		}
	}()
	return func() {
		obj._deleteSignalChan(sigChan)
	}
}

type dbusPropertyPowerOnBattery struct {
	*property.BaseObserver
	core *dbus.Object
}

func (this *dbusPropertyPowerOnBattery) SetValue(notwritable interface{}) {
	fmt.Println("com.deepin.system.Power.OnBattery is not writable")
}

func (this *dbusPropertyPowerOnBattery) Get() bool {
	v, _ := this.GetValue()
	return v.(bool)
}
func (this *dbusPropertyPowerOnBattery) GetValue() (interface{} /*bool*/, error) {
	var r dbus.Variant
	err := this.core.Call("org.freedesktop.DBus.Properties.Get", 0, "com.deepin.system.Power", "OnBattery").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool), nil
	}
	return *new(bool), err
}
func (this *dbusPropertyPowerOnBattery) GetType() reflect.Type {
	return reflect.TypeOf((*bool)(nil)).Elem()
}

type dbusPropertyPowerHasBattery struct {
	*property.BaseObserver
	core *dbus.Object
}

func (this *dbusPropertyPowerHasBattery) SetValue(notwritable interface{}) {
	fmt.Println("com.deepin.system.Power.HasBattery is not writable")
}

func (this *dbusPropertyPowerHasBattery) Get() bool {
	v, _ := this.GetValue()
	return v.(bool)
}
func (this *dbusPropertyPowerHasBattery) GetValue() (interface{} /*bool*/, error) {
	var r dbus.Variant
	err := this.core.Call("org.freedesktop.DBus.Properties.Get", 0, "com.deepin.system.Power", "HasBattery").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool), nil
	}
	return *new(bool), err
}
func (this *dbusPropertyPowerHasBattery) GetType() reflect.Type {
	return reflect.TypeOf((*bool)(nil)).Elem()
}

type dbusPropertyPowerBatteryPercentage struct {
	*property.BaseObserver
	core *dbus.Object
}

func (this *dbusPropertyPowerBatteryPercentage) SetValue(notwritable interface{}) {
	fmt.Println("com.deepin.system.Power.BatteryPercentage is not writable")
}

func (this *dbusPropertyPowerBatteryPercentage) Get() float64 {
	v, _ := this.GetValue()
	return v.(float64)
}
func (this *dbusPropertyPowerBatteryPercentage) GetValue() (interface{} /*float64*/, error) {
	var r dbus.Variant
	err := this.core.Call("org.freedesktop.DBus.Properties.Get", 0, "com.deepin.system.Power", "BatteryPercentage").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64), nil
	}
	return *new(float64), err
}
func (this *dbusPropertyPowerBatteryPercentage) GetType() reflect.Type {
	return reflect.TypeOf((*float64)(nil)).Elem()
}

type dbusPropertyPowerBatteryStatus struct {
	*property.BaseObserver
	core *dbus.Object
}

func (this *dbusPropertyPowerBatteryStatus) SetValue(notwritable interface{}) {
	fmt.Println("com.deepin.system.Power.BatteryStatus is not writable")
}

func (this *dbusPropertyPowerBatteryStatus) Get() uint32 {
	v, _ := this.GetValue()
	return v.(uint32)
}
func (this *dbusPropertyPowerBatteryStatus) GetValue() (interface{} /*uint32*/, error) {
	var r dbus.Variant
	err := this.core.Call("org.freedesktop.DBus.Properties.Get", 0, "com.deepin.system.Power", "BatteryStatus").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32), nil
	}
	return *new(uint32), err
}
func (this *dbusPropertyPowerBatteryStatus) GetType() reflect.Type {
	return reflect.TypeOf((*uint32)(nil)).Elem()
}

type dbusPropertyPowerBatteryTimeToEmpty struct {
	*property.BaseObserver
	core *dbus.Object
}

func (this *dbusPropertyPowerBatteryTimeToEmpty) SetValue(notwritable interface{}) {
	fmt.Println("com.deepin.system.Power.BatteryTimeToEmpty is not writable")
}

func (this *dbusPropertyPowerBatteryTimeToEmpty) Get() uint64 {
	v, _ := this.GetValue()
	return v.(uint64)
}
func (this *dbusPropertyPowerBatteryTimeToEmpty) GetValue() (interface{} /*uint64*/, error) {
	var r dbus.Variant
	err := this.core.Call("org.freedesktop.DBus.Properties.Get", 0, "com.deepin.system.Power", "BatteryTimeToEmpty").Store(&r)
	if err == nil && r.Signature().String() == "t" {
		return r.Value().(uint64), nil
	}
	return *new(uint64), err
}
func (this *dbusPropertyPowerBatteryTimeToEmpty) GetType() reflect.Type {
	return reflect.TypeOf((*uint64)(nil)).Elem()
}

type dbusPropertyPowerBatteryTimeToFull struct {
	*property.BaseObserver
	core *dbus.Object
}

func (this *dbusPropertyPowerBatteryTimeToFull) SetValue(notwritable interface{}) {
	fmt.Println("com.deepin.system.Power.BatteryTimeToFull is not writable")
}

func (this *dbusPropertyPowerBatteryTimeToFull) Get() uint64 {
	v, _ := this.GetValue()
	return v.(uint64)
}
func (this *dbusPropertyPowerBatteryTimeToFull) GetValue() (interface{} /*uint64*/, error) {
	var r dbus.Variant
	err := this.core.Call("org.freedesktop.DBus.Properties.Get", 0, "com.deepin.system.Power", "BatteryTimeToFull").Store(&r)
	if err == nil && r.Signature().String() == "t" {
		return r.Value().(uint64), nil
	}
	return *new(uint64), err
}
func (this *dbusPropertyPowerBatteryTimeToFull) GetType() reflect.Type {
	return reflect.TypeOf((*uint64)(nil)).Elem()
}

type dbusPropertyPowerHasLidSwitch struct {
	*property.BaseObserver
	core *dbus.Object
}

func (this *dbusPropertyPowerHasLidSwitch) SetValue(notwritable interface{}) {
	fmt.Println("com.deepin.system.Power.HasLidSwitch is not writable")
}

func (this *dbusPropertyPowerHasLidSwitch) Get() bool {
	v, _ := this.GetValue()
	return v.(bool)
}
func (this *dbusPropertyPowerHasLidSwitch) GetValue() (interface{} /*bool*/, error) {
	var r dbus.Variant
	err := this.core.Call("org.freedesktop.DBus.Properties.Get", 0, "com.deepin.system.Power", "HasLidSwitch").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool), nil
	}
	return *new(bool), err
}
func (this *dbusPropertyPowerHasLidSwitch) GetType() reflect.Type {
	return reflect.TypeOf((*bool)(nil)).Elem()
}

func NewPower(destName string, path dbus.ObjectPath) (*Power, error) {
	if !path.IsValid() {
		return nil, errors.New("The path of '" + string(path) + "' is invalid.")
	}

	core := getBus().Object(destName, path)

	obj := &Power{Path: path, DestName: destName, core: core, signals: make(map[<-chan *dbus.Signal]struct{})}

	obj.OnBattery = &dbusPropertyPowerOnBattery{&property.BaseObserver{}, core}
	obj.HasBattery = &dbusPropertyPowerHasBattery{&property.BaseObserver{}, core}
	obj.BatteryPercentage = &dbusPropertyPowerBatteryPercentage{&property.BaseObserver{}, core}
	obj.BatteryStatus = &dbusPropertyPowerBatteryStatus{&property.BaseObserver{}, core}
	obj.BatteryTimeToEmpty = &dbusPropertyPowerBatteryTimeToEmpty{&property.BaseObserver{}, core}
	obj.BatteryTimeToFull = &dbusPropertyPowerBatteryTimeToFull{&property.BaseObserver{}, core}
	obj.HasLidSwitch = &dbusPropertyPowerHasLidSwitch{&property.BaseObserver{}, core}

	dbusAddMatch("type='signal',path='" + string(path) + "',interface='org.freedesktop.DBus.Properties',sender='" + destName + "',member='PropertiesChanged'")
	dbusAddMatch("type='signal',path='" + string(path) + "',interface='com.deepin.system.Power',sender='" + destName + "',member='PropertiesChanged'")
	sigChan := obj._createSignalChan()
	go func() {
		typeString := reflect.TypeOf("")
		typeKeyValues := reflect.TypeOf(map[string]dbus.Variant{})
		typeArrayValues := reflect.TypeOf([]string{})
		for v := range sigChan {
			if v.Name == "org.freedesktop.DBus.Properties.PropertiesChanged" &&
				len(v.Body) == 3 &&
				reflect.TypeOf(v.Body[0]) == typeString &&
				reflect.TypeOf(v.Body[1]) == typeKeyValues &&
				reflect.TypeOf(v.Body[2]) == typeArrayValues &&
				v.Body[0].(string) == "com.deepin.system.Power" {
				props := v.Body[1].(map[string]dbus.Variant)
				for key, _ := range props {
					if false {
					} else if key == "OnBattery" {
						obj.OnBattery.Notify()

					} else if key == "HasBattery" {
						obj.HasBattery.Notify()

					} else if key == "BatteryPercentage" {
						obj.BatteryPercentage.Notify()

					} else if key == "BatteryStatus" {
						obj.BatteryStatus.Notify()

					} else if key == "BatteryTimeToEmpty" {
						obj.BatteryTimeToEmpty.Notify()

					} else if key == "BatteryTimeToFull" {
						obj.BatteryTimeToFull.Notify()

					} else if key == "HasLidSwitch" {
						obj.HasLidSwitch.Notify()
					}
				}
			} else if v.Name == "com.deepin.system.Power.PropertiesChanged" && len(v.Body) == 1 && reflect.TypeOf(v.Body[0]) == typeKeyValues {
				for key, _ := range v.Body[0].(map[string]dbus.Variant) {
					if false {
					} else if key == "OnBattery" {
						obj.OnBattery.Notify()

					} else if key == "HasBattery" {
						obj.HasBattery.Notify()

					} else if key == "BatteryPercentage" {
						obj.BatteryPercentage.Notify()

					} else if key == "BatteryStatus" {
						obj.BatteryStatus.Notify()

					} else if key == "BatteryTimeToEmpty" {
						obj.BatteryTimeToEmpty.Notify()

					} else if key == "BatteryTimeToFull" {
						obj.BatteryTimeToFull.Notify()

					} else if key == "HasLidSwitch" {
						obj.HasLidSwitch.Notify()
					}
				}
			}
		}
	}()

	dbusAddMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='BatteryDisplayUpdate'")

	dbusAddMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='BatteryAdded'")

	dbusAddMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='BatteryRemoved'")

	dbusAddMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='LidClosed'")

	dbusAddMatch("type='signal',path='" + string(obj.Path) + "',interface='com.deepin.system.Power',sender='" + obj.DestName + "',member='LidOpened'")

	runtime.SetFinalizer(obj, func(_obj *Power) { DestroyPower(_obj) })
	return obj, nil
}
