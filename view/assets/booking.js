// ฟังก์ชันเพื่อกำหนดวันที่ขั้นต่ำเป็นวันพรุ่งนี้
function setMinDates() {
	const startDate = document.getElementById('startDate');
	const endDate = document.getElementById('endDate');

	const today = new Date();
	today.setDate(today.getDate() + 1); // เพิ่มวันให้เป็นวันพรุ่งนี้

	const year = today.getFullYear();
	const month = String(today.getMonth() + 1).padStart(2, '0');
	const day = String(today.getDate()).padStart(2, '0');

	const minDate = `${year}-${month}-${day}`;

	startDate.setAttribute('min', minDate);
	endDate.setAttribute('min', minDate);
}

// ฟังก์ชันเพื่อสร้างตัวเลือกเวลา
function populateTimeOptions(selectId) {
	const select = document.getElementById(selectId);
	const startHour = 9;  // 09:00
	const endHour = 19;   // 19:00

	for (let hour = startHour; hour <= endHour; hour++) {
		const time = hour.toString().padStart(2, '0') + ":00";
		const option = document.createElement('option');
		option.value = time;
		option.textContent = time;
		select.appendChild(option);
	}
}

// เรียกใช้ฟังก์ชันเมื่อหน้าเว็บโหลดเสร็จ
window.onload = function() {
	setMinDates();
	populateTimeOptions('startTime');
	populateTimeOptions('endTime');
};
