package pkg

type ValidationHelper interface {
	/* internal use within spex_service pkg */

	// ForEach runs the given function on the key-value pairs kept within the store.
	// It will stop executing on next key-value pair if the previous result returned is false.
	Check(key any) bool
	// WithDefaultSettings should configure the spex config store implementer itself with some predefined settings
	SetCheckers(checkers []Checker) error
}

type ValidationHelp struct {
	checkers []Checker
}

func NewValidationHelp() *ValidationHelp {
	return &ValidationHelp{
		checkers: make([]Checker, 0),
	}
}

func (v *ValidationHelp) Check(key any) bool {
	for _, checker := range v.checkers {
		if !checker.Check(key) {
			return false
		}
	}
	return true
}

func (v *ValidationHelp) SetCheckers(checkers []Checker) error {
	v.checkers = checkers
	return nil
}

func Hi() error {
	return nil
}

type Checker interface {
	// Update should update the implementer struct itself.
	// If you only reassign the concurrency-safe attributes, you don't need sync.RWMutex
	// If you update by swapping the entire object, it is better to use sync.RWMutex
	// since it changes the underlying pointer in an unsafe manner.
	Check(config any) bool
}

func NewOption() {}

type Option struct {
	checkers []Checker
	config   any
}

func newOption() *Option {
	return &Option{
		checkers: make([]Checker, 0),
	}
}

var globalOpts = newOption()

func WithChecker(f Checker) {
	globalOpts.WithChecker(f)
}

func WithConfig(f any) {
	globalOpts.WithConfig(f)
}

func Apply(helper ValidationHelper) error {
	return globalOpts.Apply(helper)
}

func (o *Option) WithChecker(f Checker) {
	o.checkers = append(o.checkers, f)
}

func (o *Option) WithConfig(f any) {
	o.config = f
}

func (o *Option) Apply(helper ValidationHelper) error {
	// TODO: recheck config

	// set checker
	helper.SetCheckers(o.checkers)

	return nil
}

// func init() {
// 	GlobalOpts.Register()
// }
