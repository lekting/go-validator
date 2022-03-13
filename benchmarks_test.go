package validator

import "testing"

func BenchmarkStructLevelValidationSuccess(b *testing.B) {
	registerRequst := &RegisterRequest{
		UserName: "s",
	}

	validate := testScheme()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = validate.Validate(registerRequst)
	}
}

func BenchmarkStructLevelValidationSuccessParallel(b *testing.B) {
	registerRequst := &RegisterRequest{
		UserName: "s",
	}

	validate := testScheme()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = validate.Validate(registerRequst)
		}
	})
}

func BenchmarkStructLevelValidationFailure(b *testing.B) {
	registerRequst := &RegisterRequest{
		UserName: "bad value",
	}

	validate := testScheme()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = validate.Validate(registerRequst)
	}
}

func BenchmarkStructLevelValidationFailureParallel(b *testing.B) {
	registerRequst := &RegisterRequest{
		UserName: "bad value",
	}

	validate := testScheme()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = validate.Validate(registerRequst)
		}
	})
}
