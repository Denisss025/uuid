package uuid

import "testing"

func TestLibUuid(t *testing.T) {
	var uuid Uuid
	uuid.Generate()
	if uuid.IsNil() {
		t.Error("Got nil from uuid.Generate()")
	}

	uuid.Clear()
	if !uuid.IsNil() {
		t.Error("uuid should be nil after clear")
	}

	uuid.GenerateRandom()
	if uuid.IsNil() {
		t.Error("Got nil from uuid.GenerateRandom()")
	}

	var uuid2 Uuid
	uuid2.GenerateRandom()
	if uuid2.IsNil() {
		t.Error("Got nil from uuid2.GenerateRandom()")
	}

	if uuid2.CompareTo(uuid) == 0 {
		t.Error("uuid2.CompareTo(uuid) should not return 0")
	}

	if uuid.CompareTo(uuid2) == 0 {
		t.Error("uuid.CompareTo(uuid2) should not return 0")
	}

	if uuid.CompareTo(uuid) != 0 {
		t.Error("Compare to self must return 0")
	}

	uuid.CopyTo(&uuid2)
	if uuid.CompareTo(uuid2) != 0 {
		t.Error("Copy does not work")
	}

	uuidStr := uuid.String()
	guid := uuid.ToGuid()
	if uuidStr != uuid2.String() {
		t.Error("uuid and uuid2 should generate the same string")
	}

	if uuidStr != uuid.String() {
		t.Error("uuid generate the same string unless new value is " +
			"generated")
	}

	if len(uuidStr) != 32 {
		t.Error("uuid.String() should generate a string with " +
			"the size of 32 characters")
	}

	if len(guid) != 36 {
		t.Error("guid should contain 36 characters")
	}
}

func BenchmarkGenerate(b *testing.B) {
	b.StopTimer()
	var uuid Uuid
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		uuid.Generate()
		b.StopTimer()
	}
}

func BenchmarkGenRandom(b *testing.B) {
	b.StopTimer()
	var uuid Uuid
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		uuid.GenerateRandom()
		b.StopTimer()
	}
}

func BenchmarkGenTime(b *testing.B) {
	var uuid Uuid
	for n := 0; n < b.N; n++ {
		uuid.GenerateTime()
	}
}

func BenchmarkGenTimeSafe(b *testing.B) {
	var uuid Uuid
	for n := 0; n < b.N; n++ {
		uuid.GenerateTimeSafe()
	}
}

func BenchmarkClear(b *testing.B) {
	var uuid Uuid
	for n := 0; n < b.N; n++ {
		uuid.Clear()
	}
}

func BenchmarkToString(b *testing.B) {
	var uuid Uuid
	uuid.GenerateRandom()

	for n := 0; n < b.N; n++ {
		uuid.String()
	}
}

func BenchmarkToGuid(b *testing.B) {
	var uuid Uuid
	uuid.GenerateRandom()

	for n := 0; n < b.N; n++ {
		uuid.ToGuid()
	}
}
