// Copyright (c) 2014 - Max Persson <max@looplab.se>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventhorizon

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&ReflectAggregateSuite{})

type ReflectAggregateSuite struct{}

func (s *ReflectAggregateSuite) TestNewReflectAggregate(c *C) {
	var nilID UUID
	agg := NewReflectAggregate(nilID, EmptyAggregate{})
	c.Assert(agg, Not(Equals), nil)
	c.Assert(agg.id, Equals, nilID)
	c.Assert(agg.eventsLoaded, Equals, 0)
	c.Assert(agg.handler, Not(Equals), nil)
	c.Assert(agg.handler, FitsTypeOf, NewReflectEventHandler(nil, ""))
}

func (s *ReflectAggregateSuite) TestAggregateID(c *C) {
	id := NewUUID()
	agg := NewReflectAggregate(id, nil)
	result := agg.AggregateID()
	c.Assert(result, Equals, id)
}

func (s *ReflectAggregateSuite) TestApplyEvent(c *C) {
	// Apply one event.
	id := NewUUID()
	agg := NewReflectAggregate(id, nil)
	mockHandler := &MockEventHandler{
		events: make([]Event, 0),
	}
	agg.handler = mockHandler
	event1 := TestEvent{NewUUID(), "event1"}
	agg.ApplyEvent(event1)
	c.Assert(agg.eventsLoaded, Equals, 1)
	c.Assert(mockHandler.events, DeepEquals, []Event{event1})

	// Apply two events.
	agg = NewReflectAggregate(id, nil)
	mockHandler = &MockEventHandler{
		events: make([]Event, 0),
	}
	agg.handler = mockHandler
	event2 := TestEvent{NewUUID(), "event2"}
	agg.ApplyEvent(event1)
	agg.ApplyEvent(event2)
	c.Assert(agg.eventsLoaded, Equals, 2)
	c.Assert(mockHandler.events, DeepEquals, []Event{event1, event2})
}

func (s *ReflectAggregateSuite) TestApplyEvents(c *C) {
	// Apply one event.
	id := NewUUID()
	agg := NewReflectAggregate(id, nil)
	mockHandler := &MockEventHandler{
		events: make([]Event, 0),
	}
	agg.handler = mockHandler
	event1 := TestEvent{NewUUID(), "event1"}
	agg.ApplyEvents([]Event{event1})
	c.Assert(agg.eventsLoaded, Equals, 1)
	c.Assert(mockHandler.events, DeepEquals, []Event{event1})

	// Apply two events.
	agg = NewReflectAggregate(id, nil)
	mockHandler = &MockEventHandler{
		events: make([]Event, 0),
	}
	agg.handler = mockHandler
	event2 := TestEvent{NewUUID(), "event2"}
	agg.ApplyEvents([]Event{event1, event2})
	c.Assert(agg.eventsLoaded, Equals, 2)
	c.Assert(mockHandler.events, DeepEquals, []Event{event1, event2})

	// Apply no event.
	agg = NewReflectAggregate(id, nil)
	mockHandler = &MockEventHandler{
		events: make([]Event, 0),
	}
	agg.handler = mockHandler
	agg.ApplyEvents([]Event{})
	c.Assert(agg.eventsLoaded, Equals, 0)
	c.Assert(mockHandler.events, DeepEquals, []Event{})
}
